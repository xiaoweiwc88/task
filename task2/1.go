package main

import (
	"fmt"
)

func addTen(num *int) {
	*num = *num + 10
}

func main() {
	a := 10
	fmt.Println("before change", a)
	addTen(&a)
	fmt.Println("after change", a)
}
