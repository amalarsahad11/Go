package main

import "fmt"

func main() {
	x := 3
	for x < 6 {
		fmt.Println(x)
		x++
	}
	for t := 0; t <= 10; t++ {
		fmt.Println(t)
	}
}
