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
}
func (g *gateway) onRegisterGatewayResponse(msg *network.Message) {
	regResult := protocol.C2GRegisterGatewayResult{}
	err := proto.Unmarshal(msg.Body, &regResult)
	if base.CheckError(err, "unmarshal register gateway result") {
		base.LogDebug("Gateway start success")
	}
}
func (g *gateway) onRegisterGatewayUserResponse(msg *network.Message) {

}
