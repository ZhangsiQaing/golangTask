package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// var counter int
	// var w sync.Mutex
	// var wg sync.WaitGroup

	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		for j := 0; j < 1000; j++ {
	// 			w.Lock()
	// 			counter++
	// 			w.Unlock()
	// 		}
	// 		defer wg.Done()
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(counter)

	var counter int64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)

}
