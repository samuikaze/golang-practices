package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示類別 T 實作了介面 I，但在 Go 中不需要顯式宣告實作
func (t T) M() {
	fmt.Printf("t.S = %v", t.S)
}

func interfaces() {
	var i I = T{"Hello"}
	i.M()
}
