package main

import "fmt"

func isPalindrome(x int) bool {
	// 负数直接不是回文
	// 如果末尾是0，但本身不为0（例如10，100),也不是回文
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	original := x

	for x > 0 {
		// 取末尾一位平拼接reversed的末尾
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

func main() {
	array := [...]int{433, -44, 55, 121, -123, 999, 10}

	for idx, v := range array {
		if isPalindrome(v) {
			fmt.Printf("arr[%d]=%d 是回文数\n", idx, v)
		} else {
			fmt.Printf("arr[%d]=%d 不是回文数\n", idx, v)
		}
	}
	return
}
