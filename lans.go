package share

import (
	"github.com/golang/protobuf/proto"

	"util/etcd"
	"util/logs"
	"util/run"

	"core/net"
	"core/net/dispatcher/pb"
	"core/net/lan"
	"core/net/lan/pipe"

	"share/msg"
)

var _ = logs.Debug

//
var (
	g_lan    *pipe.Lan
	g_chMsg  chan *pb.PbFrame
	g_handle func(*pb.PbFrame)
)

const x_chanMsg_procNum = 10    // 处理goroutine数
const x_chanMsg_cacheNum = 1000 // chan中缓存数据量

// 服务器间相关处理
func InitLans(lanCfg *lan.LanCfg, etcdCfg *etcd.SrvCfg, handle func(*pb.PbFrame)) {
	// init
	g_lan = pipe.NewLan(lanCfg)
	g_chMsg = make(chan *pb.PbFrame, x_chanMsg_cacheNum)
	g_handle = handle

	// reg and watch
	etcd.RegAndWatchs(lanCfg.Name, etcdCfg, g_lan.Update)

	// recv msg
	go run.Exec(true, recvMsg)

	// proc msg
	procMsg()
}

// recv server msg and send it to channel
func recvMsg() {
	for {
		raw, e := g_lan.Server.Recv()
		if e != nil {
			logs.Panicln(e)
		}

		f := &pb.PbFrame{}
		if e := proto.Unmarshal(raw, f); e != nil {
			logs.Panicln(e)
		}

		select {
		case g_chMsg <- f:
			// do nothing
		default:
			NoticeTooBusy(f)
		}
	}
}

//
func NoticeTooBusy(f *pb.PbFrame) {
	//
	dstUrl := f.GetSrcUrl()

	// message
	nf := &pb.PbFrame{
		SrcUrl:  proto.String(f.GetDstUrls()[0]),
		DstUrls: []string{dstUrl},
		MsgRaw:  nil, // to do too busy
	}
	SendFrame2Server(dstUrl, nf)
}

//
func SendMsg(accId int64, srcUrl, dstUrl string, msgId msg.EMsg, m proto.Message) {
	//
	raw, e := PackMsg(msgId, m)
	if e != nil {
		logs.Error("pack msg failed! srcUrl:%v, dstUrl:%v, accId:%v, msgId:%v, msg:%+v",
			srcUrl, dstUrl, accId, msgId, m)
		return
	}

	// feedback
	nb := &pb.PbFrame{
		SrcUrl:  proto.String(srcUrl),
		DstUrls: []string{dstUrl},
		AccId:   proto.Int64(accId),
		MsgRaw:  raw,
	}
	SendFrame2Server(dstUrl, nb)
}

//
func SendFrame2Server(dstUrl string, f *pb.PbFrame) bool {
	//
	accId := f.AccId

	// server
	srvId, _, ok := net.Url2Part(dstUrl)
	if !ok {
		logs.Warn("invalid url: %v, accId:%v", dstUrl, accId)
		return false
	}

	d, e := proto.Marshal(f)
	if e != nil {
		logs.Warn("accId:%v, error:%v", accId, e)
		return false
	}

	if e := g_lan.Clients.SendMsg(srvId, d); e != nil {
		logs.Warn("send msg failed! accId:%v, error:%v", accId, e)
		return false
	}

	return true
}

//
func procMsg() {
	for i := 0; i < x_chanMsg_procNum; i++ {
		go func() {
			for m := range g_chMsg {
				g_handle(m)
			}
		}()
	}
}

//
func SrvUrl() string {
	return g_lan.SrvUrl
}

//
func SelectRandUrl(srv string) string {
	srvId := g_lan.Clients.SelectRand(srv)
	if "" == srvId {
		return ""
	}
	return net.GenUrl(srvId, "0")
}
