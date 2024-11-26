package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func multipleReturns(x, y string) (string, string) {
	return y, x
}

func returnNamedVars(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

func main() {
	var x1, y1 = 42, 13
	var x2, y2 = 43, 15
	fmt.Printf("%v + %v = %v\n", x1, y1, add(x1, y1))
	fmt.Printf("%v + %v = %v\n", x2, y2, add2(x2, y2))

	a, b := multipleReturns("hello", "world")
	fmt.Printf("a: %v, b: %v\n", a, b)

	x3, y3 := returnNamedVars(17)
	fmt.Printf("%v, %v\n", x3, y3)
}
