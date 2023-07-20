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

	for _, v := range args[0] {
		if v >= 'a' && v <= 'z' {
			for i := (v - 'a'); i >= 0; i-- {
				z01.PrintRune(v)
			}
		} else if v >= 'A' && v <= 'Z' {
			for i := (v - 'A'); i >= 0; i-- {
				z01.PrintRune(v)
			}
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
