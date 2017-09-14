package main

import "fmt"
import "time"

func main() {
	fmt.Println("outter")

	c := make(chan bool)
	go func() {
		fmt.Println("inner")
		time.Sleep(1 * 1000)
		fmt.Println("inner wakeup")
		close(c)
	}()
	<-c //主线程会因此阻塞
	fmt.Println("outter again")

}
