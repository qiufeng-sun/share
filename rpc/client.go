package rpc

import (
	"github.com/golang/protobuf/proto"

	"util/logs"
	"util/etcd"

	"core/net/lan/rpc"

	"share/handler"
)

var _= logs.Debug

//
func NewClientPool(cfg *rpc.PoolConfig, etcdAddrs []string, etcdWatch string) *rpc.ClientPool {
	pool := rpc.NewClientPool(cfg)
	etcd.AddWatchs(etcdAddrs, []string{etcdWatch}, pool.Update)
	return pool
}

//
var g_clientPool *rpc.ClientPool

func InitClient(cfg *rpc.PoolConfig, etcdAddrs []string, etcdWatch string) {
	g_clientPool = NewClientPool(cfg, etcdAddrs, etcdWatch)
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
