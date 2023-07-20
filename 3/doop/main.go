package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	if !(args[1] == "-" || args[1] == "+" || args[1] == "*" || args[1] == "/" || args[1] == "%") {
		return
	}

	arg1 := 0
	sign1 := 1
	for i, v := range args[0] {
		if v == '-' && i == 0 {
			sign1 = -sign1
			continue
		}
		if !(v >= '0' && v <= '9') {
			return
		}
		arg1 = arg1*10 + int(v-'0')*sign1
	}
	if sign1 == 1 && arg1 < 0 || sign1 == -1 && arg1 > 0 {
		return
	}

	arg2 := 0
	sign2 := 1
	for i, v := range args[2] {
		if v == '-' && i == 0 {
			sign2 = -sign2
			continue
		}
		if !(v >= '0' && v <= '9') {
			return
		}
		arg2 = arg2*10 + int(v-'0')*sign2
	}
	if sign2 == 1 && arg2 < 0 || sign2 == -1 && arg2 > 0 {
		return
	}

	Max := 9223372036854775807
	Min := -9223372036854775808
	result := 0

	if args[1] == "+" {
		if (arg1 > 0 && arg2 > 0) && (arg1 > Max-arg2 && arg2 > Max-arg1) || (arg1 < 0 && arg2 < 0) && (arg1 < Min-arg2 && arg2 < Min-arg1) {
			return
		} else {
			result = arg1 + arg2
		}
	} else if args[1] == "-" {
		if (arg1 < 0 && arg2 > 0) && (arg1 < Min+arg2 && arg2 > Max+arg1) {
			return
		} else {
			result = arg1 - arg2
		}
	} else if args[1] == "*" {
		if (arg2 > 0 && arg1 < 0) && (arg1 < Min/arg2 && arg2 > Min/arg1) || (arg2 < 0 && arg1 > 0) && (arg1 > Min/arg2 && arg2 < Min/arg1) ||
			(arg2 > 0 && arg1 > 0) && (arg1 > Max/arg2 && arg2 > Max/arg1) || (arg2 < 0 && arg1 < 0) && (arg1 < Max/arg2 && arg2 < Max/arg1) {
			return
		} else {
			result = arg1 * arg2
		}
	} else if args[1] == "/" {
		if args[2] != "0" {
			result = arg1 / arg2
		} else {
			message := "No division by 0\n"
			bytes := []byte(message)
			os.Stdout.Write(bytes)
			return
		}
	} else if args[1] == "%" {
		if args[2] != "0" {
			result = arg1 % arg2
		} else {
			message := "No modulo by 0\n"
			bytes := []byte(message)
			os.Stdout.Write(bytes)
			return
		}
	}

	message := itoa(result) + "\n"
	bytes := []byte(message)
	os.Stdout.Write(bytes)
}

func itoa(n int) string {
	str := ""
	if n == 0 {
		str += "0"
	}
	if n < 0 {
		str += "-"
		n = -n
	}

	count := 0
	for d := n; d != 0; d /= 10 {
		count++
	}

	for i := count; i > 0; i-- {
		if i == 1 {
			str += string((rune('0' + (n % 10))))
		} else {
			l := 1
			for j := i - 1; j > 0; j-- {
				l = l * 10
			}
			nb := n / l
			str += string((rune('0' + (nb % 10))))
		}
	}
	return str
}

// package main

// import (
// 	"os"
// 	"fmt"
// 	"strconv"
// )

// func validateOperator(test string) bool {
// 	op := []string{"+", "-", "*", "/", "%"}
// 	for _, res:= range op {
// 		if res == test {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main(){
// 	args := os.Args[1:]
// 	if len(args) > 3 || len(args) < 3 {
// 		fmt.Print()
// 	}else{
// 		if validateOperator(args[1]) == false {
// 			fmt.Println(0)
// 		}else{
// 			premier, _ := strconv.Atoi(args[0])
// 			second, _ := strconv.Atoi(args[2])

// 			if args[1] == "%" && second == 0 {
// 				fmt.Print("No Modulo by 0\n")
// 			}else if args[1] == "/" && second == 0 {
// 				fmt.Print("No division by 0\n")
// 			}else if args[1] == "+" {
// 				fmt.Println(premier+second)
// 			}else if args[1] == "-" {
// 				fmt.Println(premier-second)
// 			}else if args[1] == "*" {
// 				fmt.Println(premier * second)
// 			}else if args[1] == "/" {
// 				fmt.Println(premier/second)
// 			}else{
// 				fmt.Println(premier%second)
// 			}

// 		}
// 	}

// }
