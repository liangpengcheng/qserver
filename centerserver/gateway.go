package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
)

type gatewayPeer struct {
	gate    *protocol.Gateway
	peer    *network.ClientPeer
	userMap map[uint64]uint64
}

func (g *gatewayPeer) removeUser(manger *gatewayManger, uid uint64) {
	delete(g.userMap, uid)
	delete(manger.userGatewayMap, uid)
}
func (g *gatewayPeer) addUser(manger *gatewayManger, uid uint64) {
	g.userMap[uid] = uid
	manger.userGatewayMap[uid] = g
}

type gatewayManger struct {
	userGatewayMap map[uint64]*gatewayPeer
	gatewayMap     map[*network.ClientPeer]*gatewayPeer
	center         *centerServer
}

func (center *centerServer) createGatewayManger() *gatewayManger {
	gateway := gatewayManger{
		userGatewayMap: make(map[uint64]*gatewayPeer),
		gatewayMap:     make(map[*network.ClientPeer]*gatewayPeer),
		center:         center,
	}
	center.proc.AddCallback(int32(protocol.G2CRegisterGateway_ID), gateway.onRegisterGateway)
	center.proc.AddCallback(int32(protocol.G2CRegisterGatewayUser_ID), gateway.onRegisterUser)
	center.proc.AddCallback(int32(protocol.G2CRemoveGatewayUser_ID), gateway.onRemoveUser)
	return &gateway
}

// 注册网关，如果同一个网关的id存在的话，这个网关会注册失败
func (gate *gatewayManger) onRegisterGateway(msg *network.Message) {
	gateway := protocol.G2CRegisterGateway{}
	err := proto.Unmarshal(msg.Body, &gateway)
	response := protocol.C2GRegisterGatewayResult{}
	if !base.CheckError(err, "unmarshal gateway message error") {
		response.Result = "failed ,unmarshal gateway message error"
		go msg.Peer.SendMessage(&response, int32(protocol.C2GRegisterGatewayResult_ID))
		return
	}
	for _, exist := range gate.gatewayMap {
		if exist.gate.GetId() == gateway.GetGate().GetId() {
			base.LogWarn("the gateway(id:%d),exists", gateway.GetGate().GetId())
			response.Result = "failed,the id of this gateway exisits"
			go msg.Peer.SendMessage(&response, int32(protocol.C2GRegisterGatewayResult_ID))
			return
		}
	}
	gate.gatewayMap[msg.Peer] = &gatewayPeer{
		peer:    msg.Peer,
		gate:    gateway.GetGate(),
		userMap: make(map[uint64]uint64),
	}
	response.Result = "success"
	go msg.Peer.SendMessage(&response, int32(protocol.C2GRegisterGatewayResult_ID))
	base.LogInfo("gateway(id:%d,address:%s) is online", gateway.GetGate().GetId(), gateway.GetGate().GetAddress())
}

// 注册用户，用于查找用户在哪个网关
func (gate *gatewayManger) onRegisterUser(msg *network.Message) {
	user := protocol.G2CRegisterGatewayUser{}
	err := proto.Unmarshal(msg.Body, &user)
	response := protocol.C2GRegisterGatewayUserResult{}

	if !base.CheckError(err, "unmarshal gateway user message error") {
		response.Result = "failed ,unmarsha gatewayuser message error"
		go msg.Peer.SendMessage(&response, int32(protocol.C2GRegisterGatewayUserResult_ID))
		return
	}
	if g, ok := gate.userGatewayMap[user.GetUid()]; ok {
		delete(gate.userGatewayMap, user.GetUid())
		base.LogDebug("find user(%d) in another gateway(%d) delete it first", user.GetUid(), g.gate.GetId())
	}
	if g, ok := gate.gatewayMap[msg.Peer]; ok {
		g.addUser(gate, user.GetUid())
	}
}

// 删除网关用户
func (gate *gatewayManger) onRemoveUser(msg *network.Message) {
	user := protocol.G2CRemoveGatewayUser{}
	err := proto.Unmarshal(msg.Body, &user)
	if base.CheckError(err, "unmarshal remove user message error") {
		if g, ok := gate.userGatewayMap[user.GetUid()]; ok {
			g.removeUser(gate, user.GetUid())
		} else {
			base.LogWarn("can't find user gateway ")
		}
	}
}

// check peer is a gateway or not and remove it
func (gate *gatewayManger) removeGateway(peer *network.ClientPeer) bool {
	if g, ok := gate.gatewayMap[peer]; ok {
		for _, v := range g.userMap {
			delete(gate.userGatewayMap, v)
		}
		base.LogDebug("lost gateway(%d) connection", g.gate.GetId())
		delete(gate.gatewayMap, peer)
		return true
	}
	return false
}
