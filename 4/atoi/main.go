package main

import (
	"fmt"
)

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

func Atoi(s string) int {
	result := 0
	sign := 1
	for i, char := range s {
		if i == 0 && char == '+' {
			continue
		} else if i == 0 && char == '-' {
			sign = -sign
		} else {
			digit := int(char - '0')
			if digit < 0 || digit > 9 {
				return 0
			}
			result = result*10 + digit*sign
		}
	}
	return result
}
