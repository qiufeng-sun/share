// 消息处理注册管理
package handler

//
import (
	"util/logs"
	"util/run"

	"core/net/dispatcher/pb"
)

//
var g_frameHandlers = NewMsgHandler()

//
type hfunc func(*pb.PbFrame)

func (this hfunc) Handle(receiver interface{}, buf []byte) {
	this(receiver.(*pb.PbFrame))
}

func RegFunc(msgId int32, h func(*pb.PbFrame)) {
	g_frameHandlers.RegHandler(msgId, hfunc(h))
}

func HandleFrame(f *pb.PbFrame) {
	defer run.Recover(false)

	//
	id, ok := ParseMsgId(f.MsgRaw)
	if !ok {
		logs.Warnln("parse msg id failed!")
		return
	}

	h, info, ok := g_frameHandlers.Handler(id)
	if !ok {
		logs.Warn("not found msg handler! fromUrl:%v, msgId:%v", *f.SrcUrl, id)
		return
	}

	info.AddStats()
	h.Handle(f, nil)
}
