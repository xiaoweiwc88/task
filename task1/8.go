package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}

	return nil
}

func main() {
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	fmt.Printf("输入: %v, target=%d, 输出: %v\n", nums1, target1, twoSum(nums1, target1))

	nums2 := []int{3, 2, 4}
	target2 := 6
	fmt.Printf("输入: %v, target=%d, 输出: %v\n", nums2, target2, twoSum(nums2, target2))

	nums3 := []int{3, 3}
	target3 := 6
	fmt.Printf("输入: %v, target=%d, 输出: %v\n", nums3, target3, twoSum(nums3, target3))
}
