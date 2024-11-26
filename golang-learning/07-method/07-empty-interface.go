package main

import "fmt"

func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func emptyInterface() {
	var i interface{}
	describe3(i)

	i = "hello"
	describe3(i)

	s := i.(string)
	fmt.Printf("s = i.(string) = %s\n", s)

	s, ok := i.(string)
	fmt.Printf("s, ok = i.(string) = %s, %v\n", s, ok)

	f, ok := i.(float64)
	fmt.Printf("f, ok = i.(float64) = %f, %v\n", f, ok)
}
