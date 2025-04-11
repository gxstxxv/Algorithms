package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	x := rand.IntN(100)
	y := rand.IntN(100)
	start := time.Now()
	e := KaratsubaMultiplikation(x, y)
	elapsed := time.Since(start)

	fmt.Println(e)
	fmt.Println(elapsed)
}

func KaratsubaMultiplikation(x, y int) int {
	if x < 10 && y < 10 {
		return x * y
	}

	n := max(len(fmt.Sprint(x)), len(fmt.Sprint(y)))
	m := int(math.Ceil(float64(n / 2)))

	a := int(math.Floor(float64(x) / math.Pow10(m)))
	b := x % int(math.Pow10(m))

	c := int(math.Floor(float64(y) / math.Pow10(m)))
	d := y % int(math.Pow10(m))

	e1 := KaratsubaMultiplikation(a, c)
	e2 := KaratsubaMultiplikation(b, d)
	e3 := KaratsubaMultiplikation(a+b, c+d)
	e4 := e3 - e1 - e2

	return int(int(math.Pow10(m*2))*e1 + int(math.Pow10(m))*e4 + e2)
}
