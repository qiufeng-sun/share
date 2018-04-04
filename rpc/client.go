package rpc

import (
	"github.com/golang/protobuf/proto"

	"core/net/lan/rpc"
	"share/handler"
	"util/logs"
)

var _ = logs.Debug

//
var g_clientPool *rpc.ClientPool

func InitClient(cfg *rpc.PoolConfig) func(srv string, addrs []string) {
	g_clientPool = rpc.NewClientPool(cfg)

	return g_clientPool.Update
}

func GetClient() *rpc.Client {
	return g_clientPool.GetClient()
}

func ReturnClient(c *rpc.Client) {
	g_clientPool.ReturnClient(c)
}

func Call(raw []byte) ([]byte, error) {
	return g_clientPool.Call(raw)
}

func CallPbMsg(msgId int32, msg proto.Message, out proto.Message) error {
	b, e := handler.PackMsg(msgId, msg)
	if e != nil {
		return e
	}

	r, e := g_clientPool.Call(b)
	if e != nil {
		return e
	}

	e = handler.ParseMsgData(r, out)
	if e != nil {
		return e
	}

	return nil
}
