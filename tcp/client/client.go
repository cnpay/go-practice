package main

import (
	"io/ioutil"
	"net"
)

const ADDRESS = "localhost:7777"

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", ADDRESS)
	conn, _ := net.DialTCP("tcp4", nil, tcpAddr)

	result, _ := ioutil.ReadAll(conn)

	println(string(result))

}
