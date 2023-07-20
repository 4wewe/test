// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// args := os.Args[1:]
// if len(args) != 1 {
// 	return
// }
// 	str1 := ""
// 	str2 := ""
// 	t := ""
// 	ti := ""
// 	s := ""
// 	si := ""
// 	d := ""
// 	di := ""
// 	e := ""
// 	ei := ""
// 	for i, v := range args[0] {
// 		if i == 0 && v == '0' || i == 0 && v == '4' && len(args[0]) == 4 || !(v >= '0' && v <= '9') {
// 			for _, c := range "ERROR: cannot convert to roman digit" {
// 				z01.PrintRune(c)
// 			}
// 			z01.PrintRune('\n')
// 			return
// 		}
// 	}
// 	for i, v := range args[0] {
// 		if len(args[0])-i == 4 {
// 			if v == '1' {
// 				t, ti = "M", "M"
// 			} else if v == '2' {
// 				t, ti = "MM", "M+M"
// 			} else {
// 				t, ti = "MMM", "M+M+M"
// 			}
// 		} else if len(args[0])-i == 3 {
// 			if v == '1' {
// 				s, si = "C", "C"
// 			} else if v == '2' {
// 				s, si = "CC", "C+C"
// 			} else if v == '3' {
// 				s, si = "CCC", "C+C+C"
// 			} else if v == '4' {
// 				s, si = "CD", "(D-C)"
// 			} else if v == '5' {
// 				s, si = "D", "D"
// 			} else if v == '6' {
// 				s, si = "DC", "D+C"
// 			} else if v == '7' {
// 				s, si = "DCC", "D+C+C"
// 			} else if v == '8' {
// 				s, si = "DCCC", "D+C+C+C"
// 			} else if v == '9' {
// 				s, si = "CM", "(M-C)"
// 			}
// 		} else if len(args[0])-i == 2 {
// 			if v == '1' {
// 				d, di = "X", "X"
// 			} else if v == '2' {
// 				d, di = "XX", "X+X"
// 			} else if v == '3' {
// 				d, di = "XXX", "X+X+X"
// 			} else if v == '4' {
// 				d, di = "XL", "(L-X)"
// 			} else if v == '5' {
// 				d, di = "L", "L"
// 			} else if v == '6' {
// 				d, di = "LX", "L+X"
// 			} else if v == '7' {
// 				d, di = "LXX", "L+X+X"
// 			} else if v == '8' {
// 				d, di = "LXXX", "L+X+X+X"
// 			} else if v == '9' {
// 				d, di = "XC", "(C-X)"
// 			}
// 		} else if len(args[0])-i == 1 {
// 			if v == '1' {
// 				e, ei = "I", "I"
// 			} else if v == '2' {
// 				e, ei = "II", "I+I"
// 			} else if v == '3' {
// 				e, ei = "III", "I+I+I"
// 			} else if v == '4' {
// 				e, ei = "IV", "(V-I)"
// 			} else if v == '5' {
// 				e, ei = "V", "V"
// 			} else if v == '6' {
// 				e, ei = "VI", "V+I"
// 			} else if v == '7' {
// 				e, ei = "VII", "V+I+I"
// 			} else if v == '8' {
// 				e, ei = "VIII", "V+I+I+I"
// 			} else if v == '9' {
// 				e, ei = "IX", "(X-I)"
// 			}
// 		}
// 	}
// 	if len(args[0]) == 4 {
// 		if len(si) > 0 {
// 			str1 += ti + "+"
// 		} else {
// 			str1 += ti
// 		}
// 		str2 += t
// 	}
// 	if len(args[0]) >= 3 {
// 		if len(di) > 0 {
// 			str1 += si + "+"
// 		} else {
// 			str1 += si
// 		}
// 		str2 += s
// 	}
// 	if len(args[0]) >= 2 {
// 		if len(ei) > 0 {
// 			str1 += di + "+"
// 		} else {
// 			str1 += di
// 		}
// 		str2 += d
// 	}
// 	if len(args[0]) >= 1 {
// 		str1 += ei
// 		str2 += e
// 	}

// 	for _, v := range str1 {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')

//		for _, v := range str2 {
//			z01.PrintRune(v)
//		}
//		z01.PrintRune('\n')
//	}
package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	for i, v := range args[0] {
		if i == 0 && v == '0' || i == 0 && v == '4' && len(args[0]) == 4 || !(v >= '0' && v <= '9') {
			os.Stdout.WriteString("ERROR: cannot convert to roman digit" + "\n")
			return
		}
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	letters := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	letters1 := []string{"M", "(M-C)", "D", "(D-C)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	str := ""
	str1 := ""
	nbr := atoi(args[0])
	for i, v := range values {
		for nbr >= v {
			str1 += letters[i]
			str += letters1[i]
			if nbr > v {
				str += "+"
			}
			nbr -= v
		}
	}
	os.Stdout.WriteString(str + "\n" + str1 + "\n")
}

func atoi(s string) int {
	result := 0
	for _, c := range s {
		d := int(c - '0')
		result = result*10 + d
	}
	return result
}
