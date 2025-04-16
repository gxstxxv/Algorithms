package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	a := rand.Perm(200000)
	start := time.Now()
	InsertionSort(a)
	elapsed := time.Since(start)

	fmt.Println(sort.IntsAreSorted(a))
	fmt.Println(elapsed)
}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
