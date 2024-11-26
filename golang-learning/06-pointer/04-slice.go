package main

import (
	"fmt"
	"strings"
)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sliced []int = primes[1:4]

	fmt.Printf("Before sliced:\n%v\nAfter sliced (1:4):\n%v\n", primes, sliced)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	slice_a := names[0:2]
	slice_b := names[1:3]

	fmt.Printf("original array:\n%v\nslice_a:\n%v\nslice_b:\n%v\n", names, slice_a, slice_b)

	slice_b[0] = "XXX"

	fmt.Printf("slice_b now are:\n%v\n", slice_b)

	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Printf("q:\n%v\nr:\n%v\ns:\n%v\n", q, r, s)

	sliced1 := q[:2]
	sliced2 := q[2:]
	fmt.Printf("q = %v\nq[:2] = %v\nq[2:] = %v\n", q, sliced1, sliced2)
	// capacity 值為原始陣列長度，以 sliced1 為例，其 capacity 會取 q 的長度當作值輸出
	fmt.Printf("sliced1 length = %v\nsliced1 capacity = %v\n", len(sliced1), cap(sliced1))

	var sliced3 []int
	fmt.Printf("sliced3 = %v\nlen(sliced3) = %v\ncap(sliced3) = %v\n", sliced3, len(sliced3), cap(sliced3))
	if sliced3 == nil {
		fmt.Println("sliced3 is nil!")
	}

	make1 := make([]int, 5)
	make2 := make([]int, 0, 5)
	printSlice("make1", make1)
	printSlice("make2", make2)

	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var sliceToAdd []int
	sliceToAdd = append(sliceToAdd, 0)
	printSlice("sliceToAdd", sliceToAdd)
	sliceToAdd = append(sliceToAdd, 1)
	printSlice("sliceToAdd", sliceToAdd)
	sliceToAdd = append(sliceToAdd, 2, 3, 4)
	printSlice("sliceToAdd", sliceToAdd)

	forSlice := []int{2, 4, 8, 16, 32, 64, 128}
	for index, value := range forSlice {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}
