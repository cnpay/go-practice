package main

import (
	"net"
	"os"
)

func main() {
	// addr := net.ResolveUDPAddr("up4","127.0.0.1:1200")
	conn, err := net.Dial("udp", "127.0.0.1:1200")

	if nil != err {
		println(err.Error())
		os.Exit(1)
	}

	println("connected")

	conn.Write([]byte("anyting"))

	var buf [512]byte
	n, err := conn.Read(buf[0:]) //可以重复声明变量

	if nil != err {
		println("read error:", err.Error())
		os.Exit(1)
	}

	println(string(buf[0:n]))

	os.Exit(0)
}
