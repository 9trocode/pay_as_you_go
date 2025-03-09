package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	// Set up syslog logging
	sysLogger, err := syslog.New(syslog.LOG_INFO, "constant_load")
	if err != nil {
		log.Fatalf("Failed to set up syslog: %v", err)
	}
	log.SetOutput(sysLogger)
	log.Println("Starting constant load program")

	// Allocate memory to reach 512 MiB
	mem := make([]byte, 512*1024*1024)
	for i := range mem {
		mem[i] = byte(rand.Intn(256))
	}

	// Display memory usage
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Printf("Allocated memory: %d MiB\n", memStats.Alloc/(1024*1024))
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
			log.Println("Performed computation cycle")
		}
	}
}