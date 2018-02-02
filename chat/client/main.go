package main

import (
	"net"

	"github.com/liangpengcheng/qserver/protocol"

	"github.com/golang/protobuf/proto"

	"github.com/liangpengcheng/qserver/chat"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
)

func initClient() *network.ClientPeer {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if !base.CheckError(err, "dail gateway") {
		return nil
	}
	proc := network.NewProcessor()
	cli := &network.ClientPeer{
		Proc:       proc,
		Connection: conn,
	}
	proc.AddCallback(int32(chat.G2CIMMessage_ID), onChatMsg)
	proc.AddCallback(int32(chat.G2CIMChatResult_ID), onSendMsgResult)
	logintoken := protocol.C2GTokenLogin{}
	cli.SendMessage(&logintoken, int32(protocol.C2GTokenLogin_ID))
	go proc.StartProcess()
	go cli.ConnectionHandler()
	return cli
}

func onChatMsg(msg *network.Message) {
	chatmsg := chat.G2CIMMessage{}
	proto.Unmarshal(msg.Body, &chatmsg)
	base.LogInfo("%v", chatmsg)
}
func onSendMsgResult(msg *network.Message) {

	chatmsg := chat.G2CIMChatResult{}
	proto.Unmarshal(msg.Body, &chatmsg)
	base.LogInfo("%v", chatmsg)
}
func sendMessage(c *network.ClientPeer, text string) {
	chatmsg := chat.C2GIMMessage{
		Content: text,
		Channel: 1,
		Touser:  0,
	}
	c.SendMessage(&chatmsg, int32(chat.C2GIMMessage_ID))
}

func main() {
	cli := initClient()
	sendMessage(cli, "lalala")
	base.LogInfo("%v", *cli)
	c := make(chan struct{}, 1)
	<-c
	base.LogInfo("client exit")
}
