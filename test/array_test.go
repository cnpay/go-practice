package test

import "testing"

func TestArray1(t *testing.T) {
	var n [10]int
	for i := 0; i < len(n); i++ {
		n[i] = i * i
	}

	if n[3] != 9 {
		t.Error("出错了，n[3]应该是9")
	}

}

func TestArray2(t *testing.T) {
	var n = [3]int{1, 2, 3}
	for i := 0; i < len(n); i++ {
		t.Logf("n[%d]=%d", i, n[i])
	}
}
