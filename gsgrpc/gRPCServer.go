package gsgrpc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"rloop/Go-Ground-Station/gstypes"
	"rloop/Go-Ground-Station/helpers"
	"rloop/Go-Ground-Station/proto"
	"strconv"
	"sync"
	//"time"
)

type GRPCServer struct {
	serviceChan            chan<- gstypes.ServerControlWithTimeout
	commandChannel         chan<- gstypes.Command
	simCommandChannel      chan<- *gstypes.SimulatorCommandWithResponse
	simConfigChannel       chan<- *gstypes.SimulatorConfigWithResponse
	receiversChannelHolder *ChannelsHolder
	statusProvider         StatusProvider
}

type ChannelsHolder struct {
	ReceiverMutex sync.Mutex
	//struct that will prevent multiple operations on the channelholder at the same time, sort of mutex
	Coordinator gstypes.ReceiversCoordination
	//map that holds the channels to communicate with the grpc clients
	Receivers map[*chan gstypes.RealTimeDataBundle]*chan gstypes.RealTimeDataBundle
}

func (srv *GRPCServer) Ping(context.Context, *proto.Ping) (*proto.Pong, error) {
	srvStatus := srv.statusProvider.GetStatus()
	length := len(srvStatus.PortsListening)
	openPorts := make([]*proto.OpenPort, length)
	idx := 0
	for port, serving := range srvStatus.PortsListening {
		openPorts[idx] = &proto.OpenPort{
			Port:    int32(port),
			Serving: serving}
		idx++
	}
	status := &proto.ServerStatus{
		DataStoreManagerRunning: srvStatus.DataStoreManagerRunning,
		GRPCServerRunning:       srvStatus.GRPCServerRunning,
		BroadcasterRunning:      srvStatus.BroadcasterRunning,
		GSLoggerRunning:         srvStatus.GSLoggerRunning,
		OpenPorts:               openPorts}
	return &proto.Pong{Status: status}, nil
}

func (srv *GRPCServer) StreamPackets(req *proto.StreamRequest, stream proto.GroundStationService_StreamPacketsServer) error {
	var err error
	receiverChannel := make(chan gstypes.RealTimeDataBundle, 32)
	srv.addChannelToDatastoreQueue(receiverChannel)
	fmt.Println("gsgrpc channel pushed to map")
	for element := range receiverChannel {
		dataBundle := proto.DataBundle{}
		dataArray := make([]*proto.Params, len(element.Data))
		for idx := 0; idx < len(element.Data); idx++ {
			param := proto.Params{}
			param.RxTime = element.Data[idx].Data.RxTime
			param.ParamName = element.Data[idx].ParameterName
			param.PacketName = element.Data[idx].PacketName
			switch element.Data[idx].Data.ValueIndex {
			case 1:
				param.Value = &proto.Value{Index: 1, Int64Value: element.Data[idx].Data.Int64Value}
			case 2:
				param.Value = &proto.Value{Index: 2, Uint64Value: element.Data[idx].Data.Uint64Value}
			case 3:
				param.Value = &proto.Value{Index: 3, DoubleValue: element.Data[idx].Data.Float64Value}
			}
			dataArray[idx] = &param
		}
		dataBundle.Parameters = dataArray
		err = stream.Send(&dataBundle)
		if err != nil {
			srv.removeChannelFromDatastoreQueue(receiverChannel)
			break
		} else {
			fmt.Println("sent data to frontend server")
		}
	}
	return err
}

func (srv *GRPCServer) SendCommand(ctx context.Context, cmd *proto.Command) (*proto.Ack, error) {
	ack := &proto.Ack{}
	//the message that will be delivered to the pod
	var dataBytes []byte
	//any error encountered will be pushed into this variable and returned immediately to the sender
	var err error
	//fmt.Printf("Request for command: %v\n", cmd)
	node := cmd.Node
	packetType := cmd.PacketType
	data := cmd.Data
	dataLength := len(data)

	dataBytesArray := [][]byte{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	//convert the data values to bytes
	for idx := 0; idx < dataLength; idx++ {
		buf := new(bytes.Buffer)
		value := data[idx]
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			ack.Success = false
			ack.Message = err.Error()
			goto returnStatement
		} else {
			dataBytesArray[idx] = buf.Bytes()
		}
	}
	//if there's no data or not enough data populate the remaining byte slots with zero value
	for idx := dataLength; idx < 4; idx++ {
		var value int32 = 0
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, value)
		if err != nil {
			ack.Success = false
			ack.Message = err.Error()
			goto returnStatement
		} else {
			dataBytesArray[idx] = buf.Bytes()
		}
	}

	dataBytes = helpers.AppendVariadic(dataBytesArray...)
	/*
		dataBytes = append(dataBytesArray[0], dataBytesArray[1]...)
		dataBytes = append(dataBytes, dataBytesArray[2]...)
		dataBytes = append(dataBytes, dataBytesArray[3]...)
	*/

	if err == nil {
		command := gstypes.Command{
			Node:       node,
			PacketType: packetType,
			Data:       dataBytes,
		}
		srv.commandChannel <- command
		ack.Success = true
	}
returnStatement:
	return ack, err
}

func (srv *GRPCServer) ControlServer(ctx context.Context, control *proto.ServerControl) (*proto.Ack, error) {
	//ctxTimeout, _ := context.WithTimeout(ctx,time.Second*3)
	var response gstypes.Ack
	var ret *proto.Ack
	comm := make(chan gstypes.Ack)
	controlStruct := gstypes.ServerControlWithTimeout{
		Control:      gstypes.ServerControl_CommandEnum(control.Command),
		ResponseChan: comm,
		Ctx:          ctx}
	srv.serviceChan <- controlStruct

	select {
	case <-ctx.Done():
		fmt.Printf("context done: %v \n", ctx.Err())
	case response = <-comm:
		ret.Success = response.Success
		ret.Message = response.Message
	}
	return ret, nil
}

func (srv *GRPCServer) SendSimCommand(ctx context.Context, command *proto.SimCommand) (*proto.Ack, error) {
	var err error
	var ack *proto.Ack
	var req *gstypes.SimulatorCommandWithResponse
	var responseChan chan *proto.Ack

	responseChan = make(chan *proto.Ack)

	req = &gstypes.SimulatorCommandWithResponse{
		ResponseChan: responseChan,
		Command:       command}

	srv.simCommandChannel <- req
	ack = <-responseChan
	return ack, err
}
func (srv *GRPCServer) EditSimConfig(ctx context.Context, in *proto.SimParameterBundle) (*proto.Ack, error) {
	var err error
	var ack *proto.Ack
	var req *gstypes.SimulatorConfigWithResponse
	var responseChan chan *proto.Ack

	responseChan = make(chan *proto.Ack)

	req = &gstypes.SimulatorConfigWithResponse{
		ResponseChan: responseChan,
		Config:       in}

	srv.simConfigChannel <- req
	ack = <-responseChan
	return ack, err
}

func (srv *GRPCServer) addChannelToDatastoreQueue(receiverChannel chan gstypes.RealTimeDataBundle) {
	//srv.receiversChannelHolder.Coordinator.Call <- true
	//<-srv.receiversChannelHolder.Coordinator.Ack
	srv.receiversChannelHolder.ReceiverMutex.Lock()
	srv.receiversChannelHolder.Receivers[&receiverChannel] = &receiverChannel
	srv.receiversChannelHolder.ReceiverMutex.Unlock()
	//srv.receiversChannelHolder.Coordinator.Done <- true
}

func (srv *GRPCServer) removeChannelFromDatastoreQueue(receiverChannel chan gstypes.RealTimeDataBundle) {
	//srv.receiversChannelHolder.Coordinator.Call <- true
	//<-srv.receiversChannelHolder.Coordinator.Ack
	srv.receiversChannelHolder.ReceiverMutex.Lock()
	delete(srv.receiversChannelHolder.Receivers, &receiverChannel)
	fmt.Println("closing receiver channel")
	srv.receiversChannelHolder.ReceiverMutex.Unlock()
	//srv.receiversChannelHolder.Coordinator.Done <- true
}

func GetChannelsHolder() *ChannelsHolder {
	callChannel := make(chan bool)
	ackChannel := make(chan bool)
	doneChannel := make(chan bool)
	coordinator := gstypes.ReceiversCoordination{
		Call: callChannel,
		Ack:  ackChannel,
		Done: doneChannel}

	holder := &ChannelsHolder{
		Coordinator:   coordinator,
		ReceiverMutex: sync.Mutex{},
		Receivers:     make(map[*chan gstypes.RealTimeDataBundle]*chan gstypes.RealTimeDataBundle),
	}
	return holder
}

func newGroundStationGrpcServer(grpcChannelsHolder *ChannelsHolder, commandChannel chan<- gstypes.Command, simCommandChannel chan<- *gstypes.SimulatorCommandWithResponse, serviceChan chan<- gstypes.ServerControlWithTimeout, statusProvider StatusProvider) *GRPCServer {
	srv := &GRPCServer{
		receiversChannelHolder: grpcChannelsHolder,
		commandChannel:         commandChannel,
		serviceChan:            serviceChan,
		statusProvider:         statusProvider,
		simCommandChannel:      simCommandChannel}
	return srv
}

func NewGoGrpcServer(port int, grpcChannelsHolder *ChannelsHolder, commandChannel chan<- gstypes.Command, simCommandChannel chan<- *gstypes.SimulatorCommandWithResponse, serviceChan chan<- gstypes.ServerControlWithTimeout, statusProvider StatusProvider) (net.Listener, *grpc.Server, error) {
	GSserver := newGroundStationGrpcServer(grpcChannelsHolder, commandChannel, simCommandChannel, serviceChan, statusProvider)
	var err error
	var grpcServer *grpc.Server

	//initialize grpcserver
	strPort := ":" + strconv.Itoa(port)
	conn, err := net.Listen("tcp", strPort)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		grpcServer = grpc.NewServer()
		proto.RegisterGroundStationServiceServer(grpcServer, GSserver)
	}

	return conn, grpcServer, err
}

type StatusProvider interface {
	GetStatus() gstypes.ServiceStatus
}
