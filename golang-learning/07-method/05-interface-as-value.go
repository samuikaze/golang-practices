package main

import (
	"fmt"
	"math"
)

type F float64

type I2 interface {
	M2()
}

func (t *T) M2() {
	fmt.Printf("M2 -> t.S = %v\n", t.S)
}

func (f F) M2() {
	fmt.Printf("f = %v\n", f)
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceAsValue() {
	var i I2

	i = &T{"Hello"}
	describe(i)
	i.M2()

	i = F(math.Pi)
	describe(i)
	i.M2()
}
