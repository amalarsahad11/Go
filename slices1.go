package main

import "fmt"

func main() {
	primes := [6]int{2, 2, 7, 8, 9, 10}
	var s []int = primes[2:4]
	fmt.Println(s)
}
