package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		return
	}

	for _, v := range arg[0] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
