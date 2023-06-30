package main

import (
	"fmt"
	"time"
)

// create a function to send data
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i                 // Sending values into the channel
		time.Sleep(time.Second) // Delay between sends
	}
	close(ch) // Closing the channel after sending all values
}

// functions to return data
func receiveData(ch <-chan int) {
	for {
		value, ok := <-ch // Receiving values from the channel
		if !ok {
			// Channel has been closed
			fmt.Println("Channel closed")
			return
		}
		fmt.Println("Received:", value)
	}
}

func main() {
	ch := make(chan int) // Creating an unbuffered channel

	go sendData(ch) // Start a goroutine to send data into the channel
	receiveData(ch) // Receive data from the channel
	fmt.Println("Program ended")
}
