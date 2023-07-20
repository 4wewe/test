package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {
	str := ""
	for i := range []rune(s) {
		str += string(s[len(s)-i-1])
	}
	return str
}
