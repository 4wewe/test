package main

import "fmt"

func main() {
	base := 2
	value := -125
	fmt.Println(ItoaBase(value, base))
}

func ItoaBase(value, base int) string {
	sbase := "0123456789ABCDEF"
	sbase = sbase[:base]
	sign := ""
	if value < 0 {
		sign = "-"
		if value == -9223372036854775808 {
			value++
		}
		value = -value
	}

	result := ""
	for value > 0 {
		remainder := value % len(sbase)
		if value == 9223372036854775807 {
			remainder++
		}
		result = string(sbase[remainder]) + result
		value /= len(sbase)
	}
	return sign + result
}
