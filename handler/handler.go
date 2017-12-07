// 消息处理注册管理
package handler

//
import (
	"core/net/socket"
)

// 注册信息封装
type HandleInfo struct {
	CallTimes int64
}

func NewHandleInfo() *HandleInfo {
	return &HandleInfo{}
}

//
func (this *HandleInfo) AddStats() {
	this.CallTimes++
}

//
type MsgHandler struct {
	*socket.MsgHandler
}

//
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{socket.NewMsgHandler()}
}

// 注册
func (this *MsgHandler) RegHandler(msgId int32, handler socket.IHandler) {
	this.MsgHandler.RegHandler(msgId, handler, NewHandleInfo())
}

// 获取handler
func (this *MsgHandler) Handler(msgId int32) (socket.IHandler, *HandleInfo, bool) {
	if h, info, ok := this.MsgHandler.Handler(msgId); ok {
		return h, info.(*HandleInfo), true
	}

	return nil, nil, false
}
