package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	if len(args[0]) > 0 {
		s1 := []rune(args[1])
		s2 := []rune(args[2])
		for _, v := range args[0] {
			if v == s1[0] {
				v = s2[0]
			}
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
