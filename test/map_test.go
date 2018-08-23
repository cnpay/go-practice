package test

import (
	"fmt"
	"testing"
)

func Test_map1(t *testing.T) {
	iMap := make(map[string]int)
	iMap["liunix"] = 35
	iMap["jack"] = 24

	fmt.Println(iMap)
	fmt.Println(iMap["liunix"])

	for key, value := range iMap {
		fmt.Printf("key=%s,value=%d \n", key, value)
	}
}
