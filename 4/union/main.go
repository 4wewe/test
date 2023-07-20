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
	t := []rune(args[1])

	str := ""
	flag := 0

	for f, k := range t {
		count := 0
		flag++
		for i, c := range s {
			counter := 0
			for j := 0; j <= i; j++ {
				if c == s[j] {
					counter++
				}
				if k == c {
					count++
				}

			}
			if counter == 1 && flag == 1 {
				str += string(c)
			}
		}
		for j := 0; j <= f; j++ {
			if k == t[j] {
				count++
			}
		}
		if count == 1 {
			str += string(k)
		}
	}

	for _, c := range str {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
