package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

// outerS 继承 innerS
type outerS struct {
	b      int
	c      float32
	int    // anonymous field
	innerS //anonymous field
}

func main() {
	inner := innerS{5, 6}
	outer := outerS{8, 9.6, 5, innerS{11, 22}}
	fmt.Println(inner)
	fmt.Println(outer.in2)
	outer.Prin()
	fmt.Println(outer.in2)
}

// 父类的引用方法只能改变自己的成员变量，不能改变子类的成员变量
func (i *innerS) Prin() {
	i.in2 = 100
	fmt.Println("Pring Method")
}
