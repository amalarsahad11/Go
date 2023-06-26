https://github.com/amalarsahad11/GoLang.git
package main

import "fmt"

func main() {
	var data = [4]int{11, 12, 13, 14}
	for _, value := range data {
		fmt.Println(value)
	}
	var sdata = make([]int, 3)
	sdata[0]=15
	sdata[1]=16 
	sdata[2]=16
	for _, value := range sdata {}
}
