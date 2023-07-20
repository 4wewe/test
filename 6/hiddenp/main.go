package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}

	if len(args[0]) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}
	tmp := 0
	s := []rune(args[1])
	for i, v := range args[0] {
		j := tmp
		for ; j < len(s); j++ {
			if v == s[j] {
				tmp = j + 1
				if i == len(args[0])-1 {
					z01.PrintRune('1')
					z01.PrintRune('\n')
					return
				}
				break
			}
			if j == len(s)-1 && !(v == s[j]) {
				z01.PrintRune('0')
				z01.PrintRune('\n')
				return
			}
		}
	}
}
