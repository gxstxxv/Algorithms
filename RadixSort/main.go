package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const limit int = 10

func main() {
	a := rand.Perm(50000000)

	start := time.Now()
	RadixSort(a)
	elapsed := time.Since(start)

	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(elapsed)
}

func RadixSort(a []int) {
	m := getMaximum(a)
	for p := 1; m/p > 0; p *= 10 {
		countingSort(a, p)
	}
}

func countingSort(a []int, p int) {
	s := len(a)
	r := make([]int, s)
	c := [limit]int{0}

	for i := range s {
		c[(a[i]/p)%10]++
	}
	for i := 1; i < limit; i++ {
		c[i] += c[i-1]
	}
	for i := s - 1; i >= 0; i-- {
		r[c[(a[i]/p)%10]-1] = a[i]
		c[(a[i]/p)%10]--
	}
	for i := range s {
		a[i] = r[i]
	}
}

func getMaximum(a []int) int {
	m := a[0]
	for i := range a {
		if m < a[i] {
			m = a[i]
		}
	}
	return m
}
