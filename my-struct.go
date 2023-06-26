package main

import "fmt"

type User struct {
	Name   string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in golang")
	amal := User{"Amal", true, 16}
	fmt.Println(amal)
	fmt.Printf("amal details are: %+v\n", amal)
	fmt.Printf("Name is %v.", amal.Name)
	fmt.Printf("Age is %v.", amal.Age)
}
