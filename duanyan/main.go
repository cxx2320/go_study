package main

import (
	"fmt"
)

func main() {
	var v interface{}
	v = 22
	// 如果v的类似是int（t=22,ok=true）断言，否则（t=0,ok=false）
	t, ok := v.(int)
	fmt.Println(t, ok)
}
