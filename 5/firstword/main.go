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
	for _, v := range args[0] {
		if v == ' ' && len(str) != 0 {
			break
		} else if v == ' ' {
			continue
		} else {
			str += string(v)
		}
	}
	if len(str) != 0 {
		print(str)
	}
}

func print(s string) {
	for _, v := range s {
		z01.PrintRune(rune(v))
	}
	z01.PrintRune('\n')
}
