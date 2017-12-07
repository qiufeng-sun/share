package handler

import (
	"util/logs"

	"core/net/msg/protobuf"

	cmsg "core/net/msg"
)

var _ = logs.Debug

//
var g_msgParser = &protobuf.PbParser{}

func ParseMsgId(buf []byte) (int32, bool) {
	return cmsg.ParseMsgId(buf)
}

func ParseMsgData(buf []byte, out interface{}) error {
	return g_msgParser.Unmarshal(buf, out)
}

func PackMsg(msgId int32, in interface{}) ([]byte, error) {
	b1, b2, e := g_msgParser.Marshal(uint32(msgId), in)
	if e != nil {
		return nil, e
	}

	if len(b2) <= 0 {
		return b1, nil
	}

	return append(b1, b2...), nil
}
