package main

import (
	"net"
	"time"
)

func handleClient(conn *net.UDPConn) {
	println("enter handle client")
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:]) //udp只有一个连接(或者socket)，一直在其上读取，和tcp的listen、accept不同

	if nil != err {
		println(err.Error())
		return
	}

	println("read:", string(buf[0:n]))
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":1200") //这里要写成udp，书上是错的
	serverConn, _ := net.ListenUDP("udp", addr)
	println("waiting for client")
	for {
		handleClient(serverConn) //复用一个conn,在子函数的read中阻塞
	}

}
