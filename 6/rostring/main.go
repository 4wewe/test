package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	s := args[0]
	slice := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if len(str) > 0 && !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9' || s[i] == 36 || s[i] == ',' || s[i] == '.') {
			slice = append(slice, str)
			str = ""
		} else if i == len(s)-1 && len(str) > 0 {
			str += string(s[i])
			slice = append(slice, str)
		} else if s[i] == ' ' {
			continue
		} else {
			str += string(s[i])
		}
	}

	if len(slice) > 1 {
		for _, v := range slice[1:] {
			for _, c := range v {
				z01.PrintRune(c)
			}
			z01.PrintRune(' ')
		}
	}
	for _, v := range slice[0] {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
