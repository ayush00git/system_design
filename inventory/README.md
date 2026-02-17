Today, I want to see how you handle **concurrency safety** and **API design** under pressure.

### The Problem: The "Flash Sale" API

We are building the backend for an e-commerce site. We are launching a "Flash Sale" for the new iPhone 16.

* **Inventory:** We have exactly **100 units** in stock.
* **Traffic:** At 12:00 PM, **10,000 users** will hit the "Buy" button at the exact same millisecond.
* **Constraint 1:** You must **never** oversell (i.e., sell 101 items).
* **Constraint 2:** The API must be fast.

### Part 1: Design the Interface

I want you to define the API endpoint and the Go structs.

1. **Endpoint:** What does the HTTP method and URL look like?
2. **Request Body:** What JSON data do you need from the user?
3. **Response:** What does the success and failure JSON look like?

### Part 2: The Core Logic (The Go Question)

Write the **Go function** that handles the purchase.

* Assume you have a global variable `var Inventory int = 100`.
* Show me how you would implement the `BuyItem()` function to ensure we don't oversell when 10,000 goroutines call it simultaneously.

*(Don't worry about the database for now—just show me how you handle the in-memory logic in Go).*

**Go ahead.**


- The Mutex version 
```bash
git checkout 12a2e6d5f3dd68c7c427cdbc3588dff05959d23e
```
