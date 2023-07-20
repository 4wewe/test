package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	nbr1, err1 := strconv.Atoi(args[0])
	if err1 != nil {
		for _, e := range err1.Error() {
			z01.PrintRune(e)
		}
	}

	nbr2, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		for _, e := range err2.Error() {
			z01.PrintRune(e)
		}
	}
	if nbr2 >= 0 {
		for i := nbr1; i <= nbr2; i++ {
			for _, v := range strconv.Itoa(i) {
				z01.PrintRune(v)
			}
			if i != nbr2 {
				z01.PrintRune(' ')
			}
		}
	} else {
		for i := nbr1; i >= nbr2; i-- {
			for _, v := range strconv.Itoa(i) {
				z01.PrintRune(v)
			}
			if i != nbr2 {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
