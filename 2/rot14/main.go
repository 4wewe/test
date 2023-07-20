package main

import (
	"github.com/01-edu/z01"
)

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Rot14(s string) string {
	str := ""
	for _, v := range s {
		if v >= 'a' && v <= 'l' || v >= 'A' && v <= 'L' {
			str += string(v + 14)
		} else if v > 'l' && v < 'l'+15 || v > 'L' && v < 'L'+15 {
			str += string(v - 12)
		} else {
			str += string(v)
		}
	}
	return str
}
