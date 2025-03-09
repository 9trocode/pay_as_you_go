package main

import (
	"runtime"
	"sync"
	"time"
)

func cpuWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	start := time.Now()
	end := start.Add(duration)
	for time.Now().Before(end) {
		// Simulate CPU work
		_ = 1 + 1 // Simple operation to keep CPU busy
	}
}

func memoryWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	mem := make([]byte, 2*1024*1024*1024) // Allocate 2 GB of memory
	start := time.Now()
	end := start.Add(duration)
	for time.Now().Before(end) {
		// Simulate memory work by touching the memory
		for i := 0; i < len(mem); i += 4096 {
			mem[i] = 0
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	runtime.GOMAXPROCS(2) // Use 2 CPU cores
	var wg sync.WaitGroup
	duration := 10 * time.Second // Test duration

	for {
		wg.Add(2)
		go cpuWork(&wg, duration)
		go memoryWork(&wg, duration)

		wg.Wait()
	}
}