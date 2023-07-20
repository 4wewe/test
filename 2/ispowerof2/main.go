package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 0 || len(arg) > 1 {
		return
	}
	n := atoi(arg[0])
	for i := 2; i <= n; i *= 2 {
		if i == n {
			print("true")
			return
		}
	}
	print("false")
}

func atoi(s string) int {
	result := 0
	for _, v := range s {
		d := int(v - '0')
		if d < 0 || d > 9 {
			return 0
		}
		result = result*10 + d
	}
	return result
}

func print(s string) {
	for _, v := range s {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
