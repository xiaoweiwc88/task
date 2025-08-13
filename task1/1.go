package main

import "fmt"


func singleNumber(nums []int) int {
	mapCount := make(map[int]int)

	// for _, value := range nums {
	// 	if _, exist := mapCount[value]; exist {
	// 		mapCount[value]++
	// 	} else {
	// 		mapCount[value] = 1
	// 	}
	// }

	for _, value := range nums {
		mapCount[value]++
	}

	for key, value := range mapCount {
		if value == 1 {
			return key
		}
	}

	return -1
}



func main() {
	test_cast := [][]int{{2, 2, 1}, {3, 1, 2, 1, 2}, {1}, {2,2,3,3}}

	for _, value := range test_cast {
		fmt.Printf("search result = %d\n", singleNumber(value))
	}
	
}
