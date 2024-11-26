package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func chooseValueOrPointerAsParameter() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// a MyFloat 實作了 Abser
	a = f
	// a *Vertex 實作了 Abser
	a = &v

	// v 是 Vertex，不是 *Vertex，因此沒有實作 Abser
	// a = v

	fmt.Printf("a.Abs() = %v\n", a.Abs())
}
