package main

import (
	"fmt"
	"net"
	"os"
)

func main2() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage:%s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)

	if addr == nil {
		fmt.Println("the addrres is ", addr.String())
	} else {
		fmt.Println("the addrres is ", addr.String())
	}

	os.Exit(0)
}
