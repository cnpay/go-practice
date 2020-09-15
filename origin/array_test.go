package main

import (
	"fmt"
	"testing"
)

/*TestArray nothing */
func TestArray(t *testing.T) {
	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("n[%d]=%d\n", j, n[j])
	}
}
