package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	a := rand.Perm(40000)

	start := time.Now()
	a = MergeSort(a)
	elapsed := time.Since(start)

	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(elapsed)
}

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	m := len(a) / 2
	x := MergeSort(a[:m])
	y := MergeSort(a[m:])

	return merge(x, y)
}

func merge(x, y []int) []int {
	r := make([]int, 0, len(x)+len(y))
	i, j := 0, 0

	for i < len(x) && j < len(y) {
		if x[i] < y[j] {
			r = append(r, x[i])
			i++
		} else {
			r = append(r, y[j])
			j++
		}
	}

	r = append(r, x[i:]...)
	r = append(r, y[j:]...)

	return r
}
