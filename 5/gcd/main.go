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

	nbr1 := atoi(args[0])
	nbr2 := atoi(args[1])
	if nbr1 < 0 || nbr2 < 0 {
		return
	}

	big := 0
	small := 0
	if nbr1 >= nbr2 {
		big, small = nbr1, nbr2
	} else {
		big, small = nbr2, nbr1
	}

	result := 1
	for i := small; i >= 1; i-- {
		if big%i == 0 && small%i == 0 {
			result = i
			break
		}
	}

	for _, v := range itoa(result) {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	result := 0
	for _, char := range s {
		d := int(char - '0')
		result = result*10 + d
	}
	if result < 0 {
		return -1
	}
	return result
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
