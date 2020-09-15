package main

//写的时候除了并行写得等着，其它的读也得等着

import (
	"sync"
	"testing"
	"time"
)

var m2 *sync.RWMutex

func read2(i int) {
	m2.RLock()
	println(i, "read start")
	time.Sleep(1 * time.Second)
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read over")
	m2.RUnlock()
}

func write(i int) {
	m2.Lock()
	println(i, "write start")
	time.Sleep(1 * time.Second)
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write over")
	m2.Unlock()
}

func Test_RWMutext2(t *testing.T) {
	m2 = new(sync.RWMutex)

	go write(1)
	go read2(2)
	go write(3)
	go read2(4)

	time.Sleep(10 * time.Second)
	println("done")
}
