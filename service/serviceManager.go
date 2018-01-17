package service

import (
	"net"

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
func NewManager(s Service, centerAddress string) *Manager {
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
	}
	centerproc.AddCallback(int32(protocol.C2SServiceRegisterResult_ID), m.onServiceRegistered)
	m.Center.SendMessage(m.getServiceProto(), int32(protocol.S2CServiceRegister_ID))
	go centerproc.StartProcess()
	return &m
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
