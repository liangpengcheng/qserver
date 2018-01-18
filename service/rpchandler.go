package service

import (
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/liangpengcheng/qcontinuum/base"
	"github.com/liangpengcheng/qcontinuum/network"
)

// RPCMessage rpc消息
type RPCMessage struct {
	MSG   []byte
	UID   uint64
	MSGID int32
}

// MsgCallback 消息处理函数
type MsgCallback func(msg *RPCMessage) (proto.Message, int32)

// RPCHandler rpc处理器
type RPCHandler struct {
	MessageChan    chan *RPCMessage
	CallbackMap    map[int32]MsgCallback
	updateCallback network.ProcFunction
	FuncChan       chan network.ProcFunction
	// 更新时间
	loopTime time.Duration
	// ImmediateMode 立即回调消息，如果想要线程安全，必须设置为false，默认为false
	ImmediateMode bool
}

// NewHandler 新建处理器，包含初始化操作
func NewHandler() *RPCHandler {
	return NewHandlerWithLoopTime(24 * time.Hour)
}

// NewHandlerWithLoopTime 指定定时器
func NewHandlerWithLoopTime(time time.Duration) *RPCHandler {
	p := &RPCHandler{
		MessageChan:   make(chan *RPCMessage, 1024),
		FuncChan:      make(chan network.ProcFunction, 64),
		CallbackMap:   make(map[int32]MsgCallback),
		loopTime:      time,
		ImmediateMode: false,
	}
	return p
}

// AddCallback 设置回调
func (p *RPCHandler) addCallback(id int32, callback MsgCallback) {
	p.CallbackMap[id] = callback
}

// AddCallback 设置回调
func (p *RPCHandler) AddCallback(id int32, callback MsgCallback) {
	if id != 0 {
		p.addCallback(id, callback)
	}
}

// SetUpdate 设置更新时间，以及更新函数
func (p *RPCHandler) SetUpdate(uptime time.Duration, upcall network.ProcFunction) {
	p.loopTime = uptime
	p.updateCallback = upcall
}

// StartProcess 开始处理信息
// 只有调用了这个借口，处理器才会处理实际的信息，以及实际发送消息
func (p *RPCHandler) StartProcess() {
	defer func() {
		if err := recover(); err != nil {
			base.LogError("%v", err)
		}
	}()
	//go p.send()
	base.LogInfo("processor is starting ")
	tick := time.Tick(p.loopTime)
	for {
		select {
		case msg := <-p.MessageChan:
			if cb, ok := p.CallbackMap[msg.MSGID]; ok {
				cb(msg)
			} else {
				base.LogWarn("can't find callback(%d)", msg.MSGID)
			}

		case f := <-p.FuncChan:
			f()
		case <-tick:
			if p.updateCallback != nil {
				p.updateCallback()
			}
		}

	}
}
