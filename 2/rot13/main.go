package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 1 {
		return
	}

	for _, v := range args[0] {
		if v >= 'a' && v <= 'm' || v >= 'A' && v <= 'M' {
			z01.PrintRune(v + 13)
		} else if v > 'm' && v < 'm'+14 || v > 'M' && v < 'M'+14 {
			z01.PrintRune(v - 13)
		} else {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
