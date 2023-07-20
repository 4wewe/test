package main

import "github.com/01-edu/z01"

func main() {
	n := +1
	s := itoa(n)
	for _, v := range s {
		z01.PrintRune(v)
	}
}
func itoa(n int) string {
	str := ""
	if n == 0 {
		str = "0"
		return str
	}
	if n < 0 {
		str += "-"
		if n == -9223372036854775808 {
			n = -223372036854775808
			str += "9"
		}
		n = -n
	}
	count := 0
	for d := n; d != 0; d /= 10 {
		count++
	}
	for i := count; i > 0; i-- {
		if i == 1 {
			str += string(rune('0' + (n % 10)))
		} else {
			l := 1
			for j := i - 1; j > 0; j-- {
				l = l * 10
			}
			nb := n / l
			str += string(rune('0' + (nb % 10)))
		}
	}
	return str
}
