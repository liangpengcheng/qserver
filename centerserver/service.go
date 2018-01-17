package main

import (
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/liangpengcheng/qcontinuum/base"

	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
)

type service struct {
	connection *network.ClientPeer
	serv       *protocol.Service
}
type serviceManager struct {
	services sync.Map
}

func newServiceManager() *serviceManager {
	return &serviceManager{}
}
func (sm *serviceManager) regCallback(proc *network.Processor) {
	proc.AddCallback(int32(protocol.S2CServiceRegister_ID), sm.onServiceReg)
}

// 还需要向所有的gateway发送service节点信息
func (sm *serviceManager) addService(serv *service) {
	regresp := protocol.C2SServiceRegisterResult{
		Result: "Success",
	}
	serv.connection.SendMessage(&regresp, int32(protocol.C2SServiceRegisterResult_ID))
	sm.services.Store(serv.connection, serv)
}

// 先检查，在移除
func (sm *serviceManager) checkLostConnection(conn *network.ClientPeer) bool {
	isservice := false
	sm.services.Range(
		func(key, v interface{}) bool {
			connK := key.(*network.ClientPeer)
			if connK == conn {
				isservice = true
				return false
			}
			return true
		})

	if isservice {
		sm.removeService(conn)
	}
	return isservice
}

// 还需要向所有的gateway发送service节点信息
func (sm *serviceManager) removeService(conn *network.ClientPeer) {
	base.LogInfo("正在移除service")
	sm.services.Delete(conn)
}
func (sm *serviceManager) getService(conn *network.ClientPeer) *service {
	if val, ok := sm.services.Load(conn); ok {
		return val.(*service)
	}
	return nil
}

func (sm *serviceManager) onServiceReg(msg *network.Message) {
	reg := protocol.S2CServiceRegister{}
	err := proto.Unmarshal(msg.Body, &reg)
	if base.CheckError(err, "unmarshal s2cserviceregister") {
		serv := service{
			connection: msg.Peer,
			serv:       reg.GetServ(),
		}
		sm.addService(&serv)
	}
}
