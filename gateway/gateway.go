package gateway

import (
	"fmt"

	"github.com/zs2619/kissgateway/pb"
	"github.com/zs2619/kissnet-go"
	"google.golang.org/protobuf/proto"
)

func GWClientCB(conn kissnet.IConnection, msg []byte) error {
	if msg == nil {
		//退出
		userClinetMgr.DelUserClient(conn)
		return nil
	}
	if len(msg) < 2 {
		userClinetMgr.DelUserClient(conn)
		return fmt.Errorf("msg len error")
	}
	requestMsg := &pb.GWMsgRequest{}
	err := proto.Unmarshal(msg, requestMsg)
	if err != nil {
		return err
	}
	return nil
}
