package main

import "os"

func evalrpn(tks []string) {
	var flag bool
	var nums []int

	pop := func() int {
		if len(nums) == 0 {
			panic("runtime error: stack underflow")
		}
		n := len(nums) - 1
		t := nums[n]
		nums = nums[:n]
		return t
	}

	for _, tk := range tks {
		flag = false
		switch tk {
		case "+":
			if len(nums) < 2 {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			x := pop()
			y := pop()
			nums = append(nums, y+x)
		case "*":
			if len(nums) < 2 {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			x := pop()
			y := pop()
			nums = append(nums, y*x)
		case "-":
			if len(nums) < 2 {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			x := pop()
			y := pop()
			nums = append(nums, y-x)
		case "/":
			if len(nums) < 2 {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			x := pop()
			y := pop()
			if x == 0 {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			nums = append(nums, y/x)
		default:
			x := atoi(tk, &flag)
			if flag {
				message := "ERROR\n"
				bytes := []byte(message)
				os.Stdout.Write(bytes)
				return
			}
			nums = append(nums, x)
		}
	}

	if len(nums) != 1 {
		message := "ERROR\n"
		bytes := []byte(message)
		os.Stdout.Write(bytes)
		return
	}

	message := itoa(nums[0]) + "\n"
	bytes := []byte(message)
	os.Stdout.Write(bytes)
}

func main() {
	if len(os.Args) != 2 {
		message := "ERROR\n"
		bytes := []byte(message)
		os.Stdout.Write(bytes)
		return
	}
	tks := split(os.Args[1])
	evalrpn(tks)
}

func atoi(s string, flag *bool) int {
	result := 0
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(s); i++ {
		char := s[i]
		d := int(char - '0')
		if d < 0 || d > 9 {
			*flag = true
			return 0
		}
		result = result*10 + d
	}
	return result * sign
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	str := ""
	for n > 0 {
		str = string(rune('0'+(n%10))) + str
		n /= 10
	}

	return sign + str
}

func split(input string) []string {
	var tokens []string
	var token string
	for i := 0; i < len(input); i++ {
		char := input[i]
		if char == ' ' {
			if len(token) > 0 {
				tokens = append(tokens, token)
				token = ""
			}
		} else {
			token += string(char)
		}
	}
	if len(token) > 0 {
		tokens = append(tokens, token)
	}
	return tokens
}
