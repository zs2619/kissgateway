package main

import (
	"bytes"
	"flag"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/zs2619/kissgateway/pb"
	"google.golang.org/protobuf/proto"
)

var addr = flag.String("addr", "127.0.0.1:20000", "http service address")

func main() {
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	msg := &pb.GWMsgRequest{GwMsgType: pb.GWMsgEnum_world,
		Router:  "test",
		MsgData: []byte{},
	}
	msgbuff, err := proto.Marshal(msg)

	sendMsg := new(bytes.Buffer)
	sendMsg.Write(msgbuff)

	err = conn.WriteMessage(websocket.BinaryMessage, sendMsg.Bytes())
	if err != nil {
		return
	}
}
