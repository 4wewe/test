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
		for i := nbr2; i >= nbr1; i-- {
			for _, v := range strconv.Itoa(i) {
				z01.PrintRune(v)
			}
			if i != nbr1 {
				z01.PrintRune(' ')
			}
		}
	} else {
		for i := nbr2; i <= nbr1; i++ {
			for _, v := range strconv.Itoa(i) {
				z01.PrintRune(v)
			}
			if i != nbr1 {
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

// func atoi(s string) int,bool {
// 	result := 0
// 	sign := 1
// 	for i, char := range s {
// 		if i == 0 && char == '-' {
// 			sign = -sign
// 		} else {
// 			digit := int(char - '0')
// 			if digit < 0 || digit > 9 {
// 				return 0,true
// 			}
// 			result = result*10 + digit*sign
// 		}
// 	}
// 	return result,false
// }
