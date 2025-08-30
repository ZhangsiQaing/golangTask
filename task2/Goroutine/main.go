package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskScheduler struct {
	JobList []func()
	wg      *sync.WaitGroup
}

func (t *TaskScheduler) Add(f func()) {
	t.JobList = append(t.JobList, f)
}

func (t *TaskScheduler) Run() {
	for k, job := range t.JobList {
		t.wg.Add(1)
		go func(i int, f func()) {
			defer t.wg.Done()
			start := time.Now()
			f()
			fmt.Printf("任务%d,耗时：%s", i+1, time.Since(start))
		}(k, job)
	}
	t.wg.Wait()
}

func main() {

	// 题目一：
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	for i := 1; i <= 10; i += 2 {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()
	// }()
	// wg.Add(1)
	// go func() {
	// 	for i := 0; i <= 10; i += 2 {
	// 		fmt.Println(i)
	// 	}
	// 	wg.Done()
	// }()
	// wg.Wait()

	// 题目二：
	schedule := TaskScheduler{
		JobList: []func(){},
		wg:      &sync.WaitGroup{},
	}
	schedule.Add(func() {
		fmt.Println("开始任务1")
		time.Sleep(1 * time.Second)
		fmt.Println("结束任务1")
	})
	schedule.Add(func() {
		fmt.Println("开始任务2")
		time.Sleep(3 * time.Second)
		fmt.Println("结束任务2")
	})
	schedule.Add(func() {
		fmt.Println("开始任务2")
		time.Sleep(5 * time.Second)
		fmt.Println("结束任务2")
	})

	//任务开始
	fmt.Println("任务开始")
	schedule.Run()
	//任务结束
	fmt.Println("任务结束")

}
