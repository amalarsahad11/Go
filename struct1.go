package main

import "fmt"

type Adress struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	var a Adress
	fmt.Println(a)

	a1 := Adress{"Amal", "Patna", 9000765}
	fmt.Println("Adress1", a1)

	a2 := Adress{Name: "Soha", city: "Banglore", Pincode: 66543210}
	fmt.Println("Adress2", a2)
	a3 := Adress{Name: "Delhi"}
	fmt.Println("Address3:", a3)
}
