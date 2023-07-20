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
	s := args[0]
	slice := []string{}
	str := ""
	if len(s) > 0 {
		for i, v := range s {
			if i > 0 && v == ' ' && s[i-1] == ' ' {
				continue
			} else if v == ' ' && i != 0 {
				slice = append(slice, str)
				str = ""
			} else if i == len(s)-1 {
				slice = append(slice, str)
			} else {
				str += string(v)
			}
		}
	}
	if len(slice) > 0 {
		for i, v := range slice {
			for _, c := range v {
				z01.PrintRune(c)
			}
			if i != len(slice)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
