package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // 必须是 int32 或 int64 等支持原子操作的类型
	var wg sync.WaitGroup

	const (
		goroutines = 10
		incPerG    = 1000
	)

	wg.Add(goroutines)
	for g := 0; g < goroutines; g++ {
		go func() {
			defer wg.Done()
			for i := 0; i < incPerG; i++ {
				atomic.AddInt64(&counter, 1) // 原子自增
			}
		}()
	}

	wg.Wait()
	fmt.Println("final counter:", counter) // 期望输出：10000
}
