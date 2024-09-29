package main

import (
	"Minerva/socket/socket_02/proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		//n, err := reader.Read(buf[:])
		//if err == io.EOF {
		//	break
		//}
		//if err != nil {
		//	fmt.Println("read error:", err)
		//}
		//rec := string(buf[:n])
		//fmt.Println(rec)
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read msg with error:", err)
		}

		fmt.Println(msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go process(conn)
	}
}
