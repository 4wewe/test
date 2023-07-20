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

	nbr := atoi(args[0])
	if nbr < 0 {
		for _, v := range "ERROR\n" {
			z01.PrintRune(v)
		}
		return
	} else if nbr == 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	base := "0123456789abcdef"
	result := ""
	for nbr > 0 {
		remainder := nbr % len(base)
		result = string(base[remainder]) + result
		nbr /= len(base)
	}

	for _, el := range result {
		z01.PrintRune(el)
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	result := 0
	for i, char := range s {
		if i == 0 && char == '-' {
			return -1
		} else {
			d := int(char - '0')
			if d < 0 || d > 9 {
				return -1
			}
			result = result*10 + d
		}
	}
	return result
}
