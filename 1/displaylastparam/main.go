package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for _, v := range args[len(args)-1] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
