package main

import (
	"fmt"
)

func printSlice(x []int) {
	fmt.Printf("slice = %v, len =%d , cap = %d \n", x, len(x), cap(x))
}

func main() {
	var numbers []int
	printSlice(numbers)

	numbers = append(numbers, 0)
	printSlice(numbers)

	numbers = append(numbers, 1)
	printSlice(numbers)

	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers1 := make([]int, len(numbers), cap(numbers)*2)

	copy(numbers1, numbers)
	printSlice(numbers1)

}
