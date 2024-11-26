package main

import (
	"fmt"
	"image"
)

func images() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Printf("m.Bounds() = %v, m.At(0, 0).RGBA() = ", m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
