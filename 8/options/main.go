package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	validOptions := "abcdefghijklmnopqrstuvwxyz"
	slice := make([]string, 32)

	for i := range slice {
		slice[i] = "0"
	}

	if len(args) == 0 {
		os.Stdout.WriteString("options: " + validOptions + "\n")
		return
	}

	for _, arg := range args {
		for i, c := range arg {
			if arg[0:i+1] == "-h" {
				os.Stdout.WriteString("options: " + validOptions + "\n")
				return
			} else { //if arg[0] == '-'
				if i > 0 && !(c >= 'a' && c <= 'z') || len(arg) == 1 {
					os.Stdout.WriteString("Invalid Option\n")
					return
				} else {
					for i, v := range validOptions {
						if v == c {
							slice[i] = "1"
							break
						}
					}
				}
			}
		}
	}
	str := ""
	for i := len(slice) - 1; i >= 0; i-- {
		str += slice[i]
		if i == 0 {
			str += "\n"
		} else if i%8 == 0 {
			str += " "
		}
	}
	os.Stdout.WriteString(str)
}
