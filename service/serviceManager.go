package service

import (
	"net"

	"github.com/golang/protobuf/proto"

	"google.golang.org/grpc"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
)

// Manager 服务管理器
type Manager struct {
	Serv       Service
	CenterProc *network.Processor // center server 消息处理器
	Center     *network.ClientPeer
	Address    string
	RPC        *grpc.Server
}

// SendMessageTo 发送消息，可以跨服发送
func (m *Manager) SendMessageTo(to uint64, msg proto.Message, msgid int32) {
	buf, err := network.GetMessageBuffer(msg, msgid)
	if base.CheckError(err, "marshal send2user message") {
		msgmerge := protocol.X2XSendMessage2User{
			Sendto:  to,
			Content: buf,
		}
		m.Center.SendMessage(&msgmerge, int32(protocol.X2XSendMessage2User_ID))
		base.LogDebug("send message 2 gateway user")

	}
}

// GetMessageID 获得这个mamager的所有可处理的消息
func (m *Manager) GetMessageID() []int32 {
	var msgs []int32
	if m.Serv != nil {
		for k := range m.Serv.GetProcessor().CallbackMap {
			msgs = append(msgs, k)
		}
	}
	return msgs
}

// NewManager 创建管理器
// centerAddress 中心服务器地址
// rpcAddress 将要建立的rpc服务的地址，给gateway调用的
func NewManager(s Service, centerAddress string, rpcAddress string) *Manager {
	connection, err := net.Dial("tcp", centerAddress)
	base.PanicError(err, "dial center server")
	centerproc := network.NewProcessor()
	m := Manager{
		Serv:       s,
		CenterProc: centerproc,
		Center: &network.ClientPeer{
			Connection: connection,
			Proc:       centerproc,
		},
		Address: rpcAddress,
	}
	lis, err := net.Listen("tcp", rpcAddress)
	if !base.CheckError(err, "rpc listen") {
		return nil
	}
	base.LogInfo("RPC Service listened(%s)", rpcAddress)
	grpcServer := grpc.NewServer()
	protocol.RegisterGatewayCallServiceServer(grpcServer, &rpcHandler{Serv: s})
	go grpcServer.Serve(lis)
	m.RPC = grpcServer
	centerproc.AddCallback(int32(protocol.C2SServiceRegisterResult_ID), m.onServiceRegistered)
	m.Center.SendMessage(m.getServiceProto(), int32(protocol.S2CServiceRegister_ID))
	return &m
}

// BlockProc 阻塞接受
func (m *Manager) BlockProc() {
	m.CenterProc.StartProcess()
}
func (m *Manager) getServiceProto() *protocol.S2CServiceRegister {
	return &protocol.S2CServiceRegister{
		Serv: &protocol.Service{
			Handler: m.GetMessageID(),
			Version: m.Serv.GetVersion(),
			Name:    m.Serv.GetName(),
			Address: m.Address,
		},
	}
}

func (m *Manager) onServiceRegistered(msg *network.Message) {

}
