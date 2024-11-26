package main

import "fmt"

func pointer() {
	i, j := 42, 2701

	// 指標指向 i
	var p *int = &i
	// 透過指標讀取 i 值
	fmt.Printf("*p = %v\n", *p)
	// 透過指標更新 i 值
	*p = 21
	// 確認 i 值
	fmt.Printf("*p 指定為 21 後，i = %v\n", i)

	// 指標指向 j
	q := &j
	// 透過指標對 j 進行數學演算
	*q = *q / 37
	// 確認 j 值
	fmt.Printf("*q / 37 後，j = %v\n", j)
}
