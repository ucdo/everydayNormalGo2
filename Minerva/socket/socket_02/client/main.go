package main

import (
	"Minerva/socket/socket_02/proto"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "ic,hello world "
		bMsg, err := proto.Encode(msg)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bMsg))
		conn.Write(bMsg)
	}
}
