package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func functions() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Printf("hypot(5, 12) = %v\n", hypot(5, 12))

	fmt.Printf("compute(hypot) = %v\n", compute(hypot))
	fmt.Printf("compute(math.Pow) = %v\n", compute(math.Pow))

	var pos, neg func(int) int = adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("i = %d, pos(i) = %d, neg(-2*i) = %d\n", i, pos(i), neg(-2*i))
	}
}
