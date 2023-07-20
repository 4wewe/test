package main

import "fmt"

func main() {
	var oct byte = 38
	fmt.Printf("%08b ||/ %08b", oct, ReverseBits(oct))
}

func ReverseBits(b byte) byte {
	r := byte(0)
	div := b
	for i := 0; i < 8; i++ {
		r = (r * 2) + (div % 2)
		div = div / 2
	}
	return r
	// div := b / 16
	// mod := b % 16

	// return b%16*16 + b/16
}
