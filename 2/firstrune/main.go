package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	str := []rune(s)
	return str[0]
}
