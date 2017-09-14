package main

//写的时候除了并行写得等着，其它的读也得等着

import (
	"sync"
	"time"
)

var m *sync.RWMutex

func read(i int) {
	m.RLock()
	println(i, "read start")
	time.Sleep(1 * time.Second)
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read over")
	m.RUnlock()
}

func write(i int) {
	m.Lock()
	println(i, "write start")
	time.Sleep(1 * time.Second)
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write over")
	m.Unlock()
}

func main() {
	m = new(sync.RWMutex)

	go write(1)
	go read(2)
	go write(3)
	go read(4)

	time.Sleep(10 * time.Second)
	println("done")
}
