package main

import (
	"fmt"
)

// 值类型：所有像int、float、bool和string这些类型都属于值类型，
// 使用这些类型的变量直接指向存在内存中的值，
// 值类型的变量的值存储在栈中。当使用等号=将一个变量的值赋给另一个变量时，
// 如 j = i ,实际上是在内存中将 i 的值进行了拷贝。可以通过 &i 获取变量 i 的内存地址
func main() {

	// t := &T{"Cxx"}
	// Bar函数参数为T类型，而t为*T类型，需要使用*t进行解指针后，Bar才能调用
	// Bar(*t)

	// t := T{"Cxx"}
	// Foo函数参数为*T类型，而t为T类型，需要使用&t进行指针后，Foo才能调用
	// Foo(&t)

	t := T{"Cxx"}
	t.Add()
	t.Mun()
}

// T is struct
type T struct {
	Name string
}

func (this T) Add()  {
	fmt.Println("Add:")
	fmt.Println(this)
}

func (this *T) Mun()  {
	fmt.Println("Mun:")
	fmt.Println(this)
}

func Bar(t T)  {
	fmt.Println("Bar:")
	fmt.Println(t)
}

func Foo(t *T)  {
	fmt.Println("Foo:")
	fmt.Println(t)
}