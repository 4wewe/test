package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(s string, base string) int {
	if len(base) < 2 || !notValid(base) {
		return 0
	}

	result := 0
	for _, el := range s {
		for i, v := range base {
			if v == el {
				result = result*len(base) + i
			}
		}
	}
	return result
}

func notValid(base string) bool {
	for i := range base {
		for j := i + 1; j < len(base); j++ {
			if base[i] == '+' || base[i] == '-' || base[i] == base[j] {
				return false
			}
		}
	}
	return true
}
