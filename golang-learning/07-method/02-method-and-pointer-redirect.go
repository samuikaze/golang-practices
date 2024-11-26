package main

import (
	"fmt"
	"math"
)

func (v *Vertex) Scale3(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs3() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodAndPointerRedirect1() {
	v := Vertex{3, 4}
	fmt.Printf("v.Abs() = %v, AbsFunc(v) = %v\n", v.Abs(), AbsFunc(v))
	v.Scale3(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	fmt.Printf("p.Abs() = %v, AbsFunc(*p) = %v\n", p.Abs(), AbsFunc(*p))
	p.Scale3(3)
	ScaleFunc(p, 8)

	fmt.Printf("v = %v, p = %v\n", v, p)
}
