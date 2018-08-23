package test

import "testing"

/**测试方法的开头一定是: Test*/
func TestMehtod1(t *testing.T) {
	var m1 = func() { //匿名方法
		println("i am m1")
	}
	m1()
}
