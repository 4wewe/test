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
	front := ""
	back := ""
	for _, v := range args[0] {
		if !(v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U') && len(front) == 0 {
			back += string(v)
		} else {
			front += string(v)
		}
	}
	if len(front) == 0 {
		for _, c := range "No vowels" {
			z01.PrintRune(c)
		}
	} else {
		for _, v := range front + back + "ay" {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}
