package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(max(a, b))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
