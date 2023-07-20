package main

import "github.com/01-edu/z01"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

func PrintNbr(n int) {
	d := n
	a := n
	if n == 0 {
		z01.PrintRune((rune('0' + n)))
	}
	if a < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			a = a + 1
		}
		a = -a
	}

	count := 0
	for d != 0 {
		d /= 10
		count++
	}

	for i := count; i > 0; i-- {
		s := a
		if i == 1 {
			if n == -9223372036854775808 {
				z01.PrintRune(rune('0' + ((s % 10) + 1)))
			} else {
				z01.PrintRune(rune('0' + (s % 10)))
			}
		} else {
			l := 1
			for j := i - 1; j > 0; j-- {
				l = l * 10
			}
			nb := s / l
			z01.PrintRune(rune('0' + (nb % 10)))
		}
	}
}

func ForEach(f func(int), a []int) {
	for _, i := range a {
		f(i)
	}
}
