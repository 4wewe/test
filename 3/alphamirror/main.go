package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	if len(args[0]) > 0 {
		for _, v := range args[0] {
			if v >= 'a' && v <= 'z' {
				z01.PrintRune('z' - v + 'a')
			} else if v >= 'A' && v <= 'Z' {
				z01.PrintRune('Z' - v + 'A')
			} else {
				z01.PrintRune(v)
			}
		}
		z01.PrintRune('\n')
	}
}
