package main

import "fmt"

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))
	fmt.Printf("%#v\n", Slice(a, 2, 4))
	fmt.Printf("%#v\n", Slice(a, -3))
	fmt.Printf("%#v\n", Slice(a, -2, -1))
	fmt.Printf("%#v\n", Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 || len(nbrs) > 2 {
		return nil
	}

	start := nbrs[0]
	end := len(a)

	if len(nbrs) == 2 {
		end = nbrs[1]
	}

	if start < 0 {
		start = len(a) + start
	}
	if end < 0 {
		end = len(a) + end
	}
	if start < 0 || start >= len(a) || end <= start || end > len(a) {
		return nil
	}

	return a[start:end]
}
