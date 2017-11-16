// 消息处理注册管理
package share

//
import (
	"util/logs"

	"core/net/dispatcher/pb"
	"core/net/socket"

	"share/msg"
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

//
type hfunc func(*pb.PbFrame)

func (this hfunc) Handle(receiver interface{}, buf []byte) {
	this(receiver.(*pb.PbFrame))
}

//
var g_msgHandlers = NewMsgHandler()

//
func RegFunc(msgId msg.EMsg, h func(*pb.PbFrame)) {
	g_msgHandlers.RegHandler(int32(msgId), hfunc(h))
}

//
func HandleMsg(f *pb.PbFrame) {
	//
	id, ok := ParseMsgId(f.MsgRaw)
	if !ok {
		logs.Warnln("parse msg id failed!")
		return
	}

	h, info, ok := g_msgHandlers.Handler(id)
	if !ok {
		logs.Warn("not found msg handler! fromUrl:%v, msgId:%v", *f.SrcUrl, id)
		return
	}

	info.AddStats()
	h.Handle(f, nil)
}
