package gateway

import (
	"github.com/zs2619/kissgateway/common"
	"github.com/zs2619/kissnet-go"
)

type ProxyServer struct {
	name string
	addr string
	conn kissnet.IConnection
}

func ProxyServerClientCB(conn kissnet.IConnection, msg []byte) error {
	return nil
}

type ProxyServerMgr struct {
	Mgr map[string]*ProxyServer
}

var GProxyServerMgr *ProxyServerMgr = &ProxyServerMgr{
	Mgr: make(map[string]*ProxyServer),
}

func (this *ProxyServerMgr) InitProxyServer() error {
	for _, v := range common.WebConfig.ProxyServerConfList {
		for _, vv := range v.Addr {
			c, err := kissnet.TcpConnector(vv, ProxyServerClientCB)
			if err != nil {
				return err
			}
			GProxyServerMgr.Mgr[v.Name] = &ProxyServer{
				name: v.Name,
				addr: vv,
				conn: c,
			}
		}
	}
	return nil
}
