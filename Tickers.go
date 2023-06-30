package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a ticker that ticks every 1 second
	ticker := time.NewTicker(1 * time.Second)

	// Run a goroutine to perform an action every time the ticker ticks
	go func() {
		for range ticker.C {
			fmt.Println("Ticker ticked!")
		}
	}()

	fmt.Println("Waiting for tickers...")

	// Sleep for 5 seconds to observe the ticker in action
	time.Sleep(5 * time.Second)

	// Stop the ticker
	ticker.Stop()

	fmt.Println("Ticker stopped!")
}
