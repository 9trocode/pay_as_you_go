package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func cpuWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	start := time.Now()
	end := start.Add(duration / 2) // Reduce CPU work duration
	for time.Now().Before(end) {
		// Simulate CPU work
		_ = 1 + 1 // Simple operation to keep CPU busy
	}
}

func memoryWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	mem := make([]byte, 4*1024*1024*1024) // Allocate 4 GB of memory
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
	runtime.GOMAXPROCS(1) // Use 1 CPU core
	var wg sync.WaitGroup
	duration := 10 * time.Second // Test duration

	fmt.Println("Starting CPU and memory work")

	for i := 0; i < 5; i++ { // Run the test 5 times
		fmt.Printf("Iteration %d: Starting CPU and memory work\n", i+1)
		wg.Add(2)
		go cpuWork(&wg, duration)
		go memoryWork(&wg, duration)

		wg.Wait()
		fmt.Printf("Iteration %d: Completed CPU and memory work\n", i+1)
	}
	fmt.Println("All iterations completed")
}