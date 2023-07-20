package main

import "fmt"

func main() {
	var oct byte = 38

	fmt.Printf("%08b || %08b\n", oct, SwapBits(oct))
}

func SwapBits(octet byte) byte {
	return octet%16*16 + octet/16
}
