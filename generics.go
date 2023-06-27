package main

import "fmt"

func main() {
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))
}
func reverse[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)

	for i, ele := range s {
		r[l-i-1] = ele
	}
	return r
}
