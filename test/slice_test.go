package test

import (
	"fmt"
	"testing"
)

func printSlice(x []int) {
	fmt.Printf("slice = %v, len =%d , cap = %d \n", x, len(x), cap(x))
}

func Test_slice1(t *testing.T) {
	var numbers []int //动态数组就是slice
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2) //扩容自动扩大一倍，但这里append(numbers,2,3,4) 的容量确是6
	printSlice(numbers)

	numbers = append(numbers, 3, 4, 5)
	printSlice(numbers)

	numbersAnother := make([]int, len(numbers), cap(numbers)*2)

	copy(numbersAnother, numbers)
	printSlice(numbersAnother)

}
