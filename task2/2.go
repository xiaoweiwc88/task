package main

import "fmt"

func mulTwo(nums *[]int) {
	for i := range *nums {
		(*nums)[i] *= 2
	}
}

func main() {
	test := []int{2, 4, 1, 3}

	mulTwo(&test)

	fmt.Println("test is ", test)
}
