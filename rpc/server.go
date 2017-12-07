package rpc

import (
	"util/logs"
	"util/etcd"

	"core/net/lan"
	"core/net/lan/rpc"
)

var _ = logs.Debug

//
func InitServer(lanCfg *lan.LanCfg, etcdCfg *etcd.SrvCfg, goNum int, handle func(msg []byte) []byte) {
	// start rpc service
	s := rpc.NewServer(lanCfg)
	s.Init(goNum, handle)
	s.Serve()

	// register
	etcd.Register(etcdCfg)
}
