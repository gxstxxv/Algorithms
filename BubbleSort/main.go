package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	a := rand.Perm(4000)

	start := time.Now()
	BubbleSort(a)
	elapsed := time.Since(start)

	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(elapsed)
}

func BubbleSort(a []int) {
	is_sorted := false

	for !is_sorted {
		for i := range a {
			if i >= len(a)-1 {
				break
			}
			if a[i] < a[i+1] {
				is_sorted = true
			} else {
				a[i], a[i+1] = a[i+1], a[i]
				is_sorted = false
				break
			}
		}
	}
}
