package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type IPhone struct {
	Number int
}
type JPhone struct{}

func (iphone IPhone) call() {
	fmt.Println("i'm iphone,my Number = ", iphone.Number)
}

func (jphone JPhone) call() {
	fmt.Println("i'm jphone")
}

func main() {
	var phone Phone //狗日的，为什么的类型一定在后面，fuck

	phone = new(IPhone)
	//下面演示如何从接口中取出结构体实例
	iphone := phone.(*IPhone)
	iphone.Number = 2345
	iphone.call()
	phone.call()

	phone = new(JPhone)
	phone.call()
}
