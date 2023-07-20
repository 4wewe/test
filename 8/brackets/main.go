package main

import "os"

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	st := []rune{}
	open := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for _, r := range s {
		if closer, ok := open[r]; ok {
			st = append(st, closer)
			continue
		}
		l := len(st) - 1
		if l < 0 || r != st[l] {
			return false
		}
		st = st[:l]
	}
	return len(st) == 0
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	for i := 0; i < len(args); i++ {
		input := ""
		for _, v := range args[i] {
			if v == '(' || v == ')' || v == '[' || v == ']' || v == '{' || v == '}' {
				input += string(v)
			}
		}
		if isValid(input) {
			os.Stdout.Write([]byte("OK\n"))
		} else {
			os.Stdout.Write([]byte("ERROR\n"))
		}
	}
}
