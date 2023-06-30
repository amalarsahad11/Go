package main

import (
	"fmt"
	"sync"
)

// define function that represents work to be done
func doWork(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark the completion of the goroutine
	fmt.Printf("Worker %d started\n", id)
	// Perform the work here
	fmt.Printf("Worker %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5
	//add the number of goroutine to the wait group
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go doWork(i, &wg) // launch a goroutine for each worker
	}

	wg.Wait()

	fmt.Println("All workers completed")
}
