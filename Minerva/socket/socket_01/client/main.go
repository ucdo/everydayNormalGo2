package main

import (
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		conn.Write([]byte("ic,hello world "))
	}
}
