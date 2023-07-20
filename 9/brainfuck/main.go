package main

import (
	"os"
)

func brainfuck(str string) {
	memory := make([]int, 2048)
	memPtr := 0
	loopCount := 0

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '>':
			memPtr++
		case '<':
			memPtr--
		case '+':
			memory[memPtr]++
		case '-':
			memory[memPtr]--
		case '.':
			os.Stdout.Write([]byte{byte(memory[memPtr])})
		case '[':
			if memory[memPtr] == 0 {
				loopCount = 1
				for loopCount != 0 {
					i++
					if str[i] == ']' {
						loopCount--
					}
					if str[i] == '[' {
						loopCount++
					}
				}
			}
		case ']':
			if memory[memPtr] != 0 {
				loopCount = 1
				for loopCount != 0 {
					i--
					if str[i] == ']' {
						loopCount++
					}
					if str[i] == '[' {
						loopCount--
					}
				}
				i--
			}
		}
	}
}

func main() {
	if len(os.Args) == 2 {
		brainfuck(os.Args[1])
	}

}
