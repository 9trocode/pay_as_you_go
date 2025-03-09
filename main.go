package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func cpuWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	start := time.Now()
	end := start.Add(duration)

	// Target CPU usage of 50%
	for time.Now().Before(end) {
		// Work for 50% of the time
		workEnd := time.Now().Add(50 * time.Millisecond)
		for time.Now().Before(workEnd) {
			// Perform some calculations to use CPU
			for i := 0; i < 1000; i++ {
				_ = math.Sqrt(float64(i))
			}
		}
		// Sleep for 50% of the time to achieve constant 0.5 CPU usage
		time.Sleep(50 * time.Millisecond)
	}
}

func memoryWork(wg *sync.WaitGroup, duration time.Duration) {
	defer wg.Done()
	start := time.Now()
	end := start.Add(duration)
	for time.Now().Before(end) {
		// Allocate 500 MB of memory
		mem := make([]byte, 500*1024*1024)
		// Simulate memory work by touching the memory
		for i := 0; i < len(mem); i += 4096 {
			mem[i] = 0
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	duration := 3 * time.Second // Test duration

	fmt.Println("Starting CPU and memory work")

	for i := 0; i < 3; i++ { // Run the test 3 times
		fmt.Printf("Iteration %d: Starting CPU and memory work\n", i+1)
		wg.Add(2)
		go cpuWork(&wg, duration)
		go memoryWork(&wg, duration)

		wg.Wait()
		fmt.Printf("Iteration %d: Completed CPU and memory work\n", i+1)
	}
	fmt.Println("All iterations completed")
}