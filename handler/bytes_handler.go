// 消息处理注册管理
package handler

//
import (
	"util/logs"
	"util/run"
)

//
var g_bytesHandlers = NewMsgHandler()

type handleBytes func(msgId int32, raw []byte) []byte

func (this handleBytes) Handle(receiver interface{}, buf []byte) {
	// do nothing, only to impl interface
	logs.Panicln("invalid call")
}

func RegHandleBytes(msgId int32, h func(msgId int32, raw []byte) []byte) {
	g_bytesHandlers.RegHandler(msgId, handleBytes(h))
}

func HandleBytes(raw []byte) []byte {
	defer run.Recover(false)

	//
	id, ok := ParseMsgId(raw)
	if !ok {
		logs.Warnln("parse msg id failed!")
		return nil
	}

	h, info, ok := g_bytesHandlers.Handler(id)
	if !ok {
		logs.Warn("not found msg handler! msgId:%v", id)
		return nil
	}

	info.AddStats()
	if f, ok := h.(handleBytes); ok {
		return f(id, raw)
	}

	logs.Panicln("invalid register function!")
	return nil
}
