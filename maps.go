package main

import "fmt"

func main() {
	/*make new , new-only allocates = no init of memory,make-allocates nd init */
	score := make(map[string]int)
	score["hitesh"] = 89
	fmt.Println(score)

	score["catherine"] = 34
	score["ron"] = 23
	score["sam"] = 56
	score["vicky"] = 78
	fmt.Println(score)

	getronscore := score["ron"]
	fmt.Println(getronscore)

	delete(score, "vicky")

    fmt.Println(score)
	}

}
