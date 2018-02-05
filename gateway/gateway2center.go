package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
	"github.com/liangpengcheng/qserver/protocol"
)

func (g *gateway) regMsgCallback() {
	g.proc.AddCallback(int32(protocol.C2GRegisterGatewayResult_ID), g.onRegisterGatewayResponse)
	g.proc.AddCallback(int32(protocol.C2GRegisterGatewayUserResult_ID), g.onRegisterGatewayUserResponse)
	g.proc.AddCallback(int32(protocol.C2GServiceRegister_ID), g.onRegisterService)
	g.proc.AddCallback(int32(protocol.C2GServiceRemove_ID), g.onRemoveService)
	g.proc.AddCallback(int32(protocol.X2XSendMessage2User_ID), g.onSend2User)
}
func (g *gateway) onRegisterGatewayResponse(msg *network.Message) {
	regResult := protocol.C2GRegisterGatewayResult{}
	err := proto.Unmarshal(msg.Body, &regResult)
	if base.CheckError(err, "unmarshal register gateway result") {
		base.LogDebug("Gateway start success")
	}
}
func (g *gateway) onRegisterGatewayUserResponse(msg *network.Message) {
	regUserResult := protocol.C2GRegisterGatewayUserResult{}
	err := proto.Unmarshal(msg.Body, &regUserResult)
	if base.CheckError(err, "unmarshal register gatewayuser result") {
		if regUserResult.GetResult() != "Success" {
			base.LogWarn("register user failed :%s", regUserResult.GetResult())
		}
	}
}
func (g *gateway) onRegisterService(msg *network.Message) {
	regServ := protocol.C2GServiceRegister{}
	err := proto.Unmarshal(msg.Body, &regServ)
	if base.CheckError(err, "unmarshal register service") {
		// 连接rpc业务处理服务器 ,并且把这个service服务器注册到消息处理系统
		go g.addService(regServ.GetServ())
	}
}
func (g *gateway) onRemoveService(msg *network.Message) {
	remServ := protocol.C2GServiceRemove{}
	err := proto.Unmarshal(msg.Body, &remServ)
	if base.CheckError(err, "unmarshal remove service ") {
		// 注销消息处理 ，断开连接
		go g.removeService(remServ.GetServ())
	}
}

func (g *gateway) onSend2User(msg *network.Message) {
	retar := protocol.X2XSendMessage2User{}
	err := proto.Unmarshal(msg.Body, &retar)
	if base.CheckError(err, "(gateway) unmarsha x2x send 2 user") {
		if u, ok := g.userMap.Load(retar.Sendto); ok {
			c := u.(*client)
			go c.connection.SendMessageBuffer(retar.Content)
		}
	}
}
