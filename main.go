package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// Allocate memory to reach 512 MiB
	mem := make([]byte, 512*1024*1024)
	for i := range mem {
		mem[i] = byte(rand.Intn(256))
	}

	// Display memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Allocated memory: %d MiB\n", memStats.Alloc/(1024*1024))

	// Create load to use around 0.3 CPU
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			start := time.Now()
			// Perform computations for ~30ms to use ~30% CPU
			for time.Since(start) < 30*time.Millisecond {
				_ = rand.Float64() * rand.Float64() // Simple computation
			}
		}
	}
}