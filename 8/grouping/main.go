package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 || len(args[1]) == 0 {
		return
	}
	str := ""
	slice := []string{}
	for i, v := range args[0] {
		if i == 0 && v != '(' || i == len(args[0]) && i != ')' {
			return
		} else if i > 0 && v >= 'a' && v <= 'z' {
			str += string(v)
		} else if v == '|' || v == ')' && i == len(args[0])-1 {
			slice = append(slice, str)
			str = ""
		}
	}
	slice1 := []string{}
	for i, v := range args[1] {
		if len(str) != 0 && !(v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' || v == '’') {
			slice1 = append(slice1, str)
			str = ""
		} else if !(v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' || v == '’') {
			continue
		} else if i == len(args[1])-1 && len(str) != 0 {
			str += string(v)
			slice1 = append(slice1, str)
			str = ""
		} else {
			str += string(v)
		}
	}
	count := 0
	for _, v := range slice1 {
		x := 0
		for i := range v {
			for _, c := range slice {
				if i+len(c) <= len(v) && v[i:i+len(c)] == c && x < len(slice) {
					count++
					x++
					str = itoa(count) + ": " + v + "\n"
					os.Stdout.WriteString(str)
				}
			}
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	str := ""
	for n > 0 {
		str = string(rune('0'+(n%10))) + str
		n /= 10
	}

	return str
}
