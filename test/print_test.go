package test

import "testing"

func TestPrint(t *testing.T) {
	println("print输出是有效果的")

	if 8 != 5 {
		t.Errorf("测试失败了: %s", "是的")
	}
}
