package main

import "fmt"

func main() {
	m_a_p := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	fmt.Println("orginal map:", m_a_p)

	m_a_p[6] = "six"
	m_a_p[7] = "seven"
	fmt.Println("map after adding new key-value pair:\n", m_a_p)
	delete(m_a_p, 5)
	delete(m_a_p, 2)
	fmt.Println("map after deletion:", m_a_p)
}
