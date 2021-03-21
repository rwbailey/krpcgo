package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
	kpb "github.com/rwbailey/krpcgo/protos"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:50000")
	if err != nil {
		log.Fatal("e1 ", err)
	}
	defer conn.Close()

	msgProto := kpb.ConnectionRequest{
		Type:       kpb.ConnectionRequest_RPC,
		ClientName: "Richard",
	}

	data, err := proto.Marshal(&msgProto)
	if err != nil {
		log.Fatal("e2 ", err)
	}
	len := int64(len(data))
	buf := make([]byte, 1)

	_ = binary.PutVarint(buf, len)

	data = append(buf, data...)

	_, _ = conn.Write(data)

	resp := make([]byte, 4096)
	length, err := conn.Read(data)
	if err != nil {
		log.Fatal("e3 ", err)
	}

	messagePb := kpb.ConnectionResponse{}
	err = proto.Unmarshal(resp[:length], &messagePb)
	if err != nil {
		log.Fatal("e4 ", err)
	}

	fmt.Println(messagePb.GetMessage())
}
