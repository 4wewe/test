package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	r := n
	for _, v := range a {
		r = f(r, v)
	}
	for _, v := range itoa(r) {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

func itoa(n int) string {
	str := ""
	if n == 0 {
		str += "0"
	}
	if n < 0 {
		str += "-"
		n = -n
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

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}
