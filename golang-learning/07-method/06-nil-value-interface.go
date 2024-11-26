package main

import "fmt"

type I3 interface {
	M3()
}

func (t *T) M3() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Printf("M3 => t.S = %v", t.S)
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func nilValueInterface() {
	var i I3

	var t *T
	i = t

	describe2(i)
	i.M3()

	i = &T{"Hello"}
	describe2(i)
	i.M3()
}
