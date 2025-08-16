package main

import (
	"context"
	"fmt"
	"sort"
	"time"
)

// Task：任务函数签名——返回错误；需要时可读取 ctx 以支持取消/超时
type Task func(ctx context.Context) error

// Result：每个任务的执行结果
type Result struct {
	Index    int
	Duration time.Duration
	Err      error
}

// RunTasks：并发执行任务，concurrency 为并发度（worker 数）
func RunTasks(ctx context.Context, tasks []Task, concurrency int) []Result {
	if concurrency <= 0 {
		concurrency = 1
	}
	type job struct {
		idx  int
		task Task
	}

	jobs := make(chan job)
	results := make(chan Result, len(tasks))

	// 启动 worker 池
	for w := 0; w < concurrency; w++ {
		go func() {
			for j := range jobs {
				start := time.Now()
				var err error
				// 保护 panic，保证单个任务异常不影响其他任务与调度器本身
				func() {
					defer func() {
						if r := recover(); r != nil {
							err = fmt.Errorf("panic: %v", r)
						}
					}()
					// 如果上层已经取消，任务仍然会被调用一次，交给任务自行检查 ctx
					err = j.task(ctx)
				}()
				results <- Result{
					Index:    j.idx,
					Duration: time.Since(start),
					Err:      err,
				}
			}
		}()
	}

	// 投递任务
	go func() {
		for i, t := range tasks {
			select {
			case <-ctx.Done():
				// 如果整体被取消，停止再投递新任务
				close(jobs)
				return
			case jobs <- job{idx: i, task: t}:
			}
		}
		close(jobs)
	}()

	// 收集结果
	out := make([]Result, 0, len(tasks))
	for i := 0; i < len(tasks); i++ {
		out = append(out, <-results)
	}
	// 为了按输入顺序返回，做个排序（也可以直接用下标写回切片）
	sort.Slice(out, func(i, j int) bool { return out[i].Index < out[j].Index })
	return out
}

// ---- 演示 ----
func main() {
	// 设置整体超时（调度器与任务可通过 ctx 感知取消）
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel()

	// 构造示例任务
	tasks := []Task{
		func(ctx context.Context) error { // Task 0：睡 150ms
			time.Sleep(150 * time.Millisecond)
			return nil
		},
		func(ctx context.Context) error { // Task 1：忙等约 100ms
			start := time.Now()
			for time.Since(start) < 100*time.Millisecond {
				// 模拟 CPU 计算
			}
			return nil
		},
		func(ctx context.Context) error { // Task 2：触发 panic
			panic("boom")
		},
		func(ctx context.Context) error { // Task 3：支持取消/超时
			t := time.NewTimer(500 * time.Millisecond)
			defer t.Stop()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-t.C:
				return nil
			}
		},
	}

	// 并发度为 2 的执行
	results := RunTasks(ctx, tasks, 2)

	// 打印每个任务的耗时与结果
	for _, r := range results {
		status := "ok"
		if r.Err != nil {
			status = r.Err.Error()
		}
		fmt.Printf("task[%d] duration=%v status=%s\n", r.Index, r.Duration, status)
	}
}

