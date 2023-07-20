package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	str := ""

	for i, c := range s {
		if i == 0 || !((s[i-1] >= 'a' && s[i-1] <= 'z') || (s[i-1] >= 'A' && s[i-1] <= 'Z') || (s[i-1] >= '0' && s[i-1] <= '9')) && s[i-1] != 39 {
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
	return str
}
