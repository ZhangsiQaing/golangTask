package main

import (
	"fmt"
	"sync"
)

func Producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {

	// ch := make(chan int, 10)
	// wg := sync.WaitGroup{}

	// wg.Add(1)
	// go func(ch chan<- int) {
	// 	for i := 1; i <= 10; i++ {
	// 		ch <- i
	// 	}
	// 	wg.Done()
	// 	defer close(ch)
	// }(ch)

	// wg.Add(1)
	// go func(ch <-chan int) {
	// 	for v := range ch {
	// 		fmt.Println(v)
	// 	}
	// 	wg.Done()
	// }(ch)

	// wg.Wait()

	ch := make(chan int, 20)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go Producer(ch, &wg)
	wg.Add(1)
	go Consumer(ch, &wg)
	wg.Wait()

}
