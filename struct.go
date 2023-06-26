package main

//Pointers to structSS
import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Amal", 25}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
