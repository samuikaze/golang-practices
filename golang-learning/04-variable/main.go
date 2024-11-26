package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 未初始化值的變數 GO 會自動為其初始化為零值
var a int
var e string
var f bool

// 多個相同類別值可以使用以下方式宣告
var b, c int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 宣告常數
const Pi = 3.14159

// 常數經度非常精確
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	Big   = 1 << 100
	Small = Big >> 99
)

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// := 只可在函數中使用，未指定類別是因為透過此運算符號 GO 會自動推斷其類別
	d := 3

	fmt.Printf("a = %v, b = %v, c = %v, d = %v, e = \"%v\", f = %v\n", a, b, c, d, e, f)

	fmt.Printf("類別：%T 值：%v\n", ToBe, ToBe)
	fmt.Printf("類別：%T 值：%v\n", MaxInt, MaxInt)
	fmt.Printf("類別：%T 值：%v\n", z, z)

	var c = uint(math.Sqrt(float64(3*3 + 4*4)))
	fmt.Printf("轉換後 c = %v\n", c)

	fmt.Printf("常數 Pi = %v", Pi)

	fmt.Printf("Big = %v, Small = %v", needFloat(Big), Small)
}
