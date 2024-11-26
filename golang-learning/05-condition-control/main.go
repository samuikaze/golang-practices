package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	defer fmt.Println("推遲字串")

	for i := 0; i < 10; i++ {
		fmt.Printf("Sum = %v\n", i)
	}

	var j int
	for j < 10 {
		fmt.Printf("j = %v\n", j)
		break
	}

	k := 1
	for k < 100 {
		k += k
		fmt.Printf("k = %v\n", k)
	}

	fmt.Printf("pow(3, 2, 10) = %f\n", pow(3, 2, 10))
	fmt.Printf("pow(3, 3, 20) = %f\n", pow(3, 3, 20))

	fmt.Println("Go 執行的系統環境: ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd
		// plan9, windows
		fmt.Printf("%s \n", os)
	}

	fmt.Println("週六是哪天?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("今天")
	case today + 1:
		fmt.Println("明天")
	case today + 2:
		fmt.Println("後天")
	default:
		fmt.Println("很多天後")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早安")
	case t.Hour() < 17:
		fmt.Println("午安")
	default:
		fmt.Println("晚安")
	}

	fmt.Println("計算中")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("完成")
}
