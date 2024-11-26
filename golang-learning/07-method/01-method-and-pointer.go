package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale2(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodAndPointer() {
	v := Vertex{3, 4}
	fmt.Printf("v = {3, 4}, v.Abs() = %f\n", v.Abs())
	fmt.Printf("v = {3, 4}, Abs(v) = %f\n", Abs(v))

	typedFunction := MyFloat(-math.Sqrt2)
	fmt.Printf("typedFunction.Abs() = %v\n", typedFunction.Abs())

	fmt.Printf("After v.Scale(10), v.Abs() = %v\n", v.Abs())

	v2 := Vertex{3, 4}
	fmt.Printf("v2 = %v, ", v2)
	Scale2(&v2, 10)
	fmt.Printf("After Scale2(&v, 10), Abs2(v2) = %v\n", Abs2(v2))
}
