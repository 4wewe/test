package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	z01.PrintRune(rune(len(args)) + '0')
	z01.PrintRune('\n')
}
