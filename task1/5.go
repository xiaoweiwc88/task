package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		digits[i]++         // 先加1
		if digits[i] < 10 { // 无进位，直接返回
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {
	// digits1 := []int{1, 2, 3}
	// digits2 := []int{4, 1, 2, 3}
	// digits3 := []int{9}
	//
	// fmt.Println("digit = %v", plusOne(digits1))
	// fmt.Println("digit = %v", plusOne(digits2))
	// fmt.Println("digit = %v", plusOne(digits3))
	digits1 := []int{1, 2, 3}
	digits2 := []int{4, 1, 2, 3}
	digits3 := []int{9}

	fmt.Printf("digit1 = %v\n", plusOne(digits1)) // [1 2 4]
	fmt.Printf("digit2 = %v\n", plusOne(digits2)) // [4 1 2 4]
	fmt.Printf("digit3 = %v\n", plusOne(digits3)) // [1 0]
}
