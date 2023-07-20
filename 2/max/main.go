package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
