package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

// LongEnough 传入的l必须是Lener的实现
func LongEnough(l Lener) {
	fmt.Println("LongEnough OK", l)
}

func main() {
	Zheng2()
}

func Zheng() {
	//证明：类型 T 的可调用方法集包含接受者为 T 的所有方法
	//证明：类型 T 的可调用方法集不包含接受者为 *T 的方法

	// 值类型
	// var lst List

	// 值类型
	lst := make(List, 0)

	// 因为lst实现了Appender的*Append方法，不是实现了Appender的Append方法，运行下一行代码会出错
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List 不实现 Appender (Append method 具有指针接收器)

	// CountInto(lst, 1, 10)

	LongEnough(lst)
}
func Zheng1() {

	//证明：类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集

	// 指针类型
	plst := &List{}

	// 指针类型
	// plst := new(List)

	CountInto(plst, 1, 10)
	LongEnough(plst)
}

func Zheng2() {
	// 值类型调用 没有指针接收器的方法Len()
	// lst := make(List, 0)
	// var lst List
	// fmt.Println(lst.Len())

	// 值类型调用 有指针接收器的方法Append()
	// lst.Append(1)
	// lst.Append(100)
	// fmt.Println(lst.Len())

	// 指针类型调用 没有指针接收器的方法Len()
	// plst := &List{}
	// plst := new(List)
	// fmt.Println(plst.Len())

	// 指针类型调用 有指针接收器的方法Append()
	// plst.Append(1)
	// plst.Append(100)
	// fmt.Println(plst.Len())
}
