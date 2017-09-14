package main

//测试多‘线程’随便读----，就是不存在啊
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

func main() {
	m = new(sync.RWMutex)

	go read(1)
	go read(2)

	time.Sleep(3 * time.Second)
}
