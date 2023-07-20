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
	if nbr > 255 || nbr < -255 {
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	base := "01"
	result := ""
	for nbr > 0 {
		remainder := nbr % len(base)
		result = string(base[remainder]) + result
		nbr /= len(base)
	}
	for i := len(result); i < 8; i++ {
		z01.PrintRune('0')
	}
	for _, el := range result {
		z01.PrintRune(el)
	}
}

func atoi(s string) int {
	result := 0
	sign := 1
	for i, v := range s {
		if i == 0 && v == '-' {
			sign = -sign
			continue
		}
		d := int(v - '0')
		if d < 0 || d > 9 {
			return 0
		}
		result = result*10 + d*sign
	}
	return result
}
