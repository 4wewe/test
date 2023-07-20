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
	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune(i + '0'))
		for _, v := range " x " {
			z01.PrintRune(v)
		}
		for _, v := range args[0] {
			z01.PrintRune(v)
		}
		for _, v := range " = " {
			z01.PrintRune(v)
		}
		for _, v := range itoa(i * atoi(args[0])) {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func atoi(s string) int {
	result := 0
	for _, v := range s {
		result = result*10 + int(v-'0')
	}
	return result
}
func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	str := ""
	count := 0
	for d := n; d > 0; d /= 10 {
		count++
	}

	for ; count > 0; count-- {
		if count == 1 {
			str += string((rune('0' + (n % 10))))
		} else {
			l := 1
			for j := count - 1; j > 0; j-- {
				l = l * 10
			}
			nb := n / l
			str += string((rune('0' + (nb % 10))))
		}
	}
	return str
}
