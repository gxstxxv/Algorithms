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
	QuickSort(a)
	elapsed := time.Since(start)

	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(elapsed)
}

func QuickSort(a []int) {
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, l, r int) {
	if l >= r {
		return
	}
	i := choosePivot(l, r)
	if i != l {
		a[l], a[i] = a[i], a[l]
	}
	j := partition(a, l, r)
	quicksort(a, l, j-1)
	quicksort(a, j+1, r)
}

func partition(a []int, l, r int) int {
	p := a[l]
	i := l + 1

	for j := l + 1; j <= r && i <= r; j++ {
		if a[j] < p {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}

	a[l], a[i-1] = a[i-1], a[l]
	return i - 1
}

func choosePivot(l, r int) int {
	return l + (r-l)/2
}
