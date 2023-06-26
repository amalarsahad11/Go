package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	mynumber := 23
	var ptr = &mynumber
	fmt.Println("Value of an actual pointer", ptr)
	fmt.Println("Value of an actual pointer", *ptr)
}
