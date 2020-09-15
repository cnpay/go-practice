package main

//测试多‘线程’随便读----，就是不存在啊
import (
	"sync"
	"testing"
	"time"
)

var m1 *sync.RWMutex

func read(i int) {

	m1.RLock()
	println(i, "read start")
	time.Sleep(1 * time.Second)
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read over")
	m1.RUnlock()

}

func Test_RWMutext1(t *testing.T) {
	m1 = new(sync.RWMutex)

	go read(1)
	go read(2)

	time.Sleep(3 * time.Second)
}
