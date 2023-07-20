package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	str := ""
	for _, s := range args {
		for i, c := range s {
			if i == len(s)-1 || !((s[i+1] >= 'a' && s[i+1] <= 'z') || (s[i+1] >= 'A' && s[i+1] <= 'Z') || (s[i+1] >= '0' && s[i+1] <= '9')) && s[i+1] != 39 {
				if c >= 'a' && c <= 'z' {
					str += string(c - 32)
				} else {
					str += string(c)
				}
			} else if c >= 'A' && c <= 'Z' {
				str += string(c + 32)
			} else {
				str += string(c)
			}
		}
		str += "\n"
	}
	for _, v := range str {
		z01.PrintRune(v)
	}
}
