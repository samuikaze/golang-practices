package main

import "fmt"

func vertex() {
	type Vertex struct {
		X int
		Y int
	}

	fmt.Printf("Vertex{1, 2} = %v\n", Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4

	fmt.Printf("v.X = %v\n", v.X)

	vertexPointer := &v
	vertexPointer.X = 1e9
	fmt.Printf("v = %v\n", v)

	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	vertexPointer2 := &Vertex{1, 2}

	fmt.Printf("v1 = %v\nv2 = %v\nv3 = %v\nvertexPointer2 = %v\n", v1, v2, v3, vertexPointer2)
}
