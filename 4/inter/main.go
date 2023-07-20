package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	s := []rune(args[0])

	str := ""
	flag := false
	for i, c := range s {
		counter := 0
		for j := 0; j <= i; j++ {
			if c == s[j] {
				counter++
			}
		}
		for _, k := range args[1] {
			if c == k {
				flag = true
			}
		}
		if counter == 1 && flag {
			str += string(c)
		}
		flag = false
	}

	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
