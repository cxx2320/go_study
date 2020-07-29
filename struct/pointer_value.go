package main

import (
	"fmt"
)

type B struct {
	thing int
}

// 加 * 修改属性会作用到外部
func (b *B) change() {
	b.thing = 11
}

func (b B) write() string {
	b.thing = 100
	return fmt.Sprint(b)
}

func main() {

	// var b1 B // b1是值
	// b1.change()
	// fmt.Println(b1.thing)
	// fmt.Println(b1.write())
	// fmt.Println(b1.thing)

	// b2 := new(B) // b2是指针
	// b2.change()
	// fmt.Println(b2.thing)
	// fmt.Println(b2.write())
	// fmt.Println(b2.thing)
}
