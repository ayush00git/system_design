package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// CONFIGURATION
	url := "http://localhost:8080/api/order"
	totalRequests := 10000 // 10000 users trying to buy
	
	// COUNTERS (Thread-safe)
	var successCount int64
	var failCount int64
	var errorCount int64

	fmt.Printf("🔥 STARTING FLASH SALE: %d Users vs 100 iPhones...\n", totalRequests)
	
	// We use a WaitGroup to wait for all 10000 goroutines to finish
	var wg sync.WaitGroup
	wg.Add(totalRequests)

	startTime := time.Now()

	for i := 0; i < totalRequests; i++ {
		go func(id int) {
			defer wg.Done()

			// The JSON Body
			jsonBody := []byte(`{
				"userName": "User",
				"address": "123 Street",
				"productName": "iPhone 16"
			}`)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
			
			if err != nil {
				atomic.AddInt64(&errorCount, 1)
				fmt.Printf("❌ Connection Error: %v\n", err)
				return
			}
			defer resp.Body.Close()

			// Check Status Code
			if resp.StatusCode == http.StatusAccepted { // 202 Accepted
				atomic.AddInt64(&successCount, 1)
			} else if resp.StatusCode == http.StatusConflict { // 409 Conflict (Sold Out)
				atomic.AddInt64(&failCount, 1)
			} else {
				atomic.AddInt64(&errorCount, 1)
				fmt.Printf("⚠️ Unexpected Status: %d\n", resp.StatusCode)
			}
		}(i)
	}

	// Wait for everyone to finish
	wg.Wait()
	duration := time.Since(startTime)

	fmt.Println("\n------------------------------------------------")
	fmt.Printf("✅ SALE OVER in %v\n", duration)
	fmt.Println("------------------------------------------------")
	fmt.Printf("📦 Successful Orders (Sold): %d\n", successCount)
	fmt.Printf("😢 Failed Orders (Sold Out): %d\n", failCount)
	fmt.Printf("❌ Errors: %d\n", errorCount)
	fmt.Println("------------------------------------------------")

	if successCount == 100 {
		fmt.Println("🎉 TEST PASSED: Exactly 100 items sold!")
	} else if successCount > 100 {
		fmt.Println("🚨 TEST FAILED: You oversold! Race condition detected.")
	} else {
		fmt.Println("🤔 TEST STRANGE: You sold less than 100. Maybe DB errors?")
	}
}