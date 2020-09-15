package main

import (
	"net"
	"reflect"
	"time"
)

func main() {
	pconn, _ := net.ListenPacket("udp", ":1234") //返回一个PacketConn，问题是这是个接口，对应的结构呢
	var buf [512]byte

	println(reflect.TypeOf(pconn).String())

	// conn = pconn.(*net.UDPConn)

	for {
		n, addr, _ := pconn.ReadFrom(buf[0:])

		println("read:", string(buf[0:n]))

		daytime := time.Now().String()
		pconn.WriteTo([]byte(daytime), addr)
	}

}
