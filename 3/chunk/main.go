package main

import (
	"fmt"
	// "github.com/01-edu/z01"
)

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	chunk := make([][]int, (len(slice)+size-1)/size)
	for i := range chunk {
		if len(slice)/(size*(i+1)) > 0 {
			chunk[i] = make([]int, size)
		} else {
			chunk[i] = make([]int, len(slice)-size*i)
		}
	}
	k := 0
	l := 0
	if len(slice) > 0 {
		for _, v := range slice {
			chunk[k][l] = v
			l++
			if l == size {
				k++
				l = 0
			}
		}
	}
	fmt.Println(chunk)
}

// func Chunk(slice []int, size int) {
// 	if size > 0 {
// 		z01.PrintRune('[')
// 		if len(slice) > 0 {
// 			count := 0
// 			for i, v := range slice {
// 				if count == 0 {
// 					z01.PrintRune('[')
// 				}
// 				z01.PrintRune(rune(v + '0'))
// 				count++
// 				if !(count == size || i == len(slice)-1) {
// 					z01.PrintRune(' ')
// 				}
// 				if count == size || i == len(slice)-1 {
// 					z01.PrintRune(']')
// 					count = 0
// 					if !(i == len(slice)-1) {
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 		}
// 		z01.PrintRune(']')
// 	}
// 	z01.PrintRune('\n')
// }
