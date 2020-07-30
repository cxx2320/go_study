package main

import "fmt"

func main() {
	checkObj()
}

// 任意类型单个值
// func Ano(value interface{}) {
// 	fmt.Println(value)
// }

// Ano 任意类型多个值
func Ano(value ...interface{}) {
	fmt.Println(value)
}

// type-switch
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

// Stringer 一个接口
type Stringer interface {
	String() string
}

// 判断变量是否实现了一个接口
func checkIn() {
	// 变量.(接口名称)
	// if sv, ok := v.(Stringer); ok {
	// 	fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	// }
}

type User struct {
	Name string
}
type Member struct {
	Name string
}

// 判断变量是否是一个结构体的实例
func checkObj() {
	// 必须通过这种方式进行实例化对象
	var user interface{}
	user = User{"Cxx"}

	// 判断变量是否是一个结构体的实例
	if _, ok := user.(User); ok {
		fmt.Println(user)
	} else {
		fmt.Println("not inter")
	}

}
