package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 3)
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	arr[0] = a
	arr[1] = b
	arr[2] = c
	sort.Ints(arr)
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
	fmt.Printf("\n%d\n%d\n%d\n", a, b, c)
}
