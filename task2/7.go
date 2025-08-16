package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int) // 无缓冲通道：发送与接收需要同步配合
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者：发送 1..10，发送完关闭通道
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch) // 告知接收方“没有更多数据了”
	}()

	// 消费者：从通道持续读取直到被关闭
	go func() {
		defer wg.Done()
		for v := range ch { // range 会在通道关闭且数据读完后退出循环
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
