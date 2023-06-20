package main

import (
	"fmt"
	"sync"
	"time"
)

const MAX = 100000

func main() {
	//countLoop()
	countGoroutine()
	//countGoroutineWithLock()
}

func printResult(res int64, start time.Time) {
	fmt.Printf("Result: %d, done in %d ns\n\n", res, time.Since(start).Nanoseconds())
}

func countLoop() {
	fmt.Println("Count in loop")
	var c int64
	start := time.Now()
	for i := 0; i < MAX; i++ {
		c++
	}
	printResult(c, start)
}

type Counter struct {
	mu    sync.Mutex
	count int64
}

func (c *Counter) Inc() {
	// c.mu.Lock()
	// defer c.mu.Unlock()
	c.count++
}

func countGoroutine() {
	fmt.Println("Count in goroutine")
	var c = &Counter{}
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < MAX; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	printResult(c.count, start)
}

func countGoroutineWithLock() {
	fmt.Println("Count in goroutine with mutex")
	var (
		c int64
	)
	mu := &sync.Mutex{}
	start := time.Now()
	for i := 0; i < MAX; i++ {
		go func() {
			mu.Lock()

			c++
			mu.Unlock()
		}()
	}

	printResult(c, start)
}
