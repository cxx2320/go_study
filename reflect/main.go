package main

import (
	"fmt"
	"reflect"
)

func main() {
	Test()
}

func Test() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value1:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	_, ok := v.Interface().(float64)
	fmt.Println(ok)
}

func Test1() {
	type Myint int
	var m Myint = 5
	v := reflect.ValueOf(m)

	// v的类型
	fmt.Println(v.Kind())

	// v的值
	fmt.Println(v.Interface())
}
