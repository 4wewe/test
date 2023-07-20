package main

import "os"

func PrintMemory(arr [10]byte) {
	for i, v := range arr {
		os.Stdout.Write([]byte{hexDigit(v >> 4), hexDigit(v)})
		os.Stdout.Write([]byte(" "))
		if (i+1)%4 == 0 {
			os.Stdout.Write([]byte("\n"))
		}
	}
	os.Stdout.Write([]byte("\n"))

	for _, v := range arr {
		if v >= 32 && v <= 126 {
			os.Stdout.Write([]byte{v})
		} else {
			os.Stdout.Write([]byte("."))
		}
	}
	os.Stdout.Write([]byte("\n"))
}

func hexDigit(b byte) byte {
	b &= 0x0F
	if b < 10 {
		return b + '0'
	}
	return b - 10 + 'a'
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
