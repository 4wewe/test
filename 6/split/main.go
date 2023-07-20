package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	slice := []string{}
	str := ""
	for i := 0; i < len(s); i++ {
		if len(s[i:]) >= len(sep) && s[i:i+len(sep)] == sep {
			slice = append(slice, str)
			str = ""
			i = i + len(sep) - 1
		} else if i == len(s)-1 {
			str += string(s[i])
			slice = append(slice, str)
		} else {
			str += string(s[i])
		}
	}
	return slice
}
