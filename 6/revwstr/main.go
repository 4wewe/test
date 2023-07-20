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
	slice := []string{}
	str := ""
	for i, v := range args[0] {
		str += string(v)
		if v == ' ' {
			slice = append(slice, str)
			str = ""
		} else if i == len(args[0])-1 {
			slice = append(slice, str)
		}
	}

	if len(slice) > 0 {
		for i := len(slice) - 1; i >= 0; i-- {
			for _, v := range slice[i] {
				z01.PrintRune(v)
			}
		}
	}
	z01.PrintRune('\n')
}
