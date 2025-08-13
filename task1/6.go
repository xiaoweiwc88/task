package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums) 
	if n == 0 {
		return 0
	}

	slow := 0

	for fast := 1;  fast < n; fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func main() {
	nums := []int{9, 9, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)

	fmt.Printf("k = %d\n", k)
	fmt.Println("nums = ", nums[:k])
}
