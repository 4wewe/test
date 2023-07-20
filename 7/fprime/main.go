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
	nbr, err := atoi(args[0])
	if err {
		return
	}
	for i := 2; nbr/i >= 1; {
		if nbr%i != 0 {
			i++
		} else {
			for _, v := range itoa(i) {
				z01.PrintRune(v)
			}
			if nbr/i > 1 {
				z01.PrintRune('*')
			}
			nbr /= i
		}
	}
}

func atoi(s string) (int, bool) {
	result := 0
	for _, char := range s {
		d := int(char - '0')
		if d < 0 || d > 9 {
			return 0, true
		}
		result = result*10 + d
	}
	return result, false
}

func itoa(n int) string {
	str := ""
	count := 0
	for d := n; d != 0; d /= 10 {
		count++
	}
	for i := count; i > 0; i-- {
		if i == 1 {
			str += string((rune('0' + (n % 10))))
		} else {
			l := 1
			for j := i - 1; j > 0; j-- {
				l = l * 10
			}
			nb := n / l
			str += string((rune('0' + (nb % 10))))
		}
	}
	return str
}
