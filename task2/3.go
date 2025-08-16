package main

import  (
	"fmt"
	"time"
)

func printOdd() {
	for i := 1; i <= 10; i+= 2 {
		fmt.Println("odd:", i)
		time.Sleep(50*time.Millisecond)
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
		time.Sleep(50*time.Millisecond)
	}
}

func main() {
	go printEven()
	go printOdd()

	time.Sleep(1 * time.Second)
}
