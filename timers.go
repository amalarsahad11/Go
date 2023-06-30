package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a timer that expires after 2 seconds
	timer := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for the timer to expire...")

	// Block until the timer expires
	<-timer.C

	fmt.Println("Timer expired!")
}
