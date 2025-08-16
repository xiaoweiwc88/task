package main

import (
	"fmt"
	"sync"
)

func main() {
	//创建一个容量为10的缓冲通道
	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	//生产者:发送100个整数
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch) //发送完毕后关闭通道
	}()


	//消费者:持续接收直到通道关闭
	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
