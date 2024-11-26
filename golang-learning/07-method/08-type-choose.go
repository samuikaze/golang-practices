package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("二倍的 %v 是 %v\n", v, v*2)
	case string:
		fmt.Printf("%q 長度是 %v 位元組\n", v, len(v))
	default:
		fmt.Printf("我不知道類型 %T\n", v)
	}
}

func typeChoose() {
	do(21)
	do("hello")
	do(true)
}
