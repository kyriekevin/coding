package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		n = -n
	}

	return n
}

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(abs(a))
}
