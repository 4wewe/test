package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str := ""

	if len(args) != 1 {
		return
	}
	v := args[0]
	for i := len(v) - 1; i >= 0; i-- {
		if v[i] == ' ' && len(str) != 0 {
			break
		} else if v[i] == ' ' {
			continue
		} else {
			str += string(v[i])
		}
	}
	if len(str) != 0 {
		print(str)
	}
}

func print(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune('\n')
}
