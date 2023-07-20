package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	result := 0
	for i := 2; i <= atoi(args[0]); i++ {
		if isPrime(i) {
			result += i
		}
	}
	for _, v := range itoa(result) {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')

}
func atoi(s string) int {
	result := 0
	for i, char := range s {
		if i == 0 && char == '-' {
			return 0
		} else {
			d := int(char - '0')
			if d < 0 || d > 9 {
				return 0
			}
			result = result*10 + d
		}
	}
	return result
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}

	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func itoa(n int) string {
	str := ""
	if n == 0 {
		str += "0"
	}
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
