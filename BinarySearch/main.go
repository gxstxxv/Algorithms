package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)

	var i int
	fmt.Print(": ")
	fmt.Scan(&i)

	start := time.Now()
	b := BinarySearch(a, i)
	elapsed := time.Since(start)
	fmt.Println(b)
	fmt.Println(elapsed)
}

func BinarySearch(d []int, v int) bool {
	if len(d) == 0 {
		return false
	}

	m := len(d) / 2

	if v == d[m] {
		return true
	}

	if v < d[m] {
		return BinarySearch(d[:m], v)
	} else {
		return BinarySearch(d[m+1:], v)
	}
}
