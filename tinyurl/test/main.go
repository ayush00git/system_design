package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// CONFIGURATION
const (
	TotalRequests = 10000   // Start with 10k. 1 Million will take a while!
	Concurrency   = 100     // How many "users" are active at once
	URL           = "http://localhost:8080/api/v1/tinyurl"
)

func main() {
	fmt.Printf("🚀 Starting stress test: %d requests with %d concurrent workers...\n", TotalRequests, Concurrency)

	var wg sync.WaitGroup
	start := time.Now()

	// Channel to send "jobs" to workers
	jobs := make(chan int, TotalRequests)

	// 1. Start Workers
	for w := 0; w < Concurrency; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// 2. Send Jobs (Requests) to the channel
	for j := 0; j < TotalRequests; j++ {
		jobs <- j
	}
	close(jobs) // Close channel so workers know when to stop

	// 3. Wait for all workers to finish
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("\n------------------------------------------------")
	fmt.Printf("✅ Test Complete!\n")
	fmt.Printf("⏱️  Time Taken: %s\n", elapsed)
	fmt.Printf("⚡ Requests Per Second: %.2f req/s\n", float64(TotalRequests)/elapsed.Seconds())
	fmt.Println("------------------------------------------------")
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	jsonBody := []byte(`{"longUrl": "https://google.com"}`)

	for range jobs {
		req, _ := http.NewRequest("POST", URL, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("❌ Network Error: %v\n", err)
			continue
		}

		// --- NEW CHECK ---
		if resp.StatusCode != 200 {
			// Read the error message from the server
			buf := new(bytes.Buffer)
			buf.ReadFrom(resp.Body)
			fmt.Printf("⚠️ Server Error (Status %d): %s\n", resp.StatusCode, buf.String())
			resp.Body.Close()
			return // STOP the worker so we don't spam errors
		}
		// -----------------

		resp.Body.Close()
	}
}