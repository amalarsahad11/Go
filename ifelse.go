package main

import "fmt"

func main() {
	var a int = 9
	if a < 20 {
		fmt.Printf("a is less than 20\n")
	} else {
		fmt.Printf("a is not less than 20\n")
	}
	fmt.Printf("value of a is : %d\n", a)
	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}
}
