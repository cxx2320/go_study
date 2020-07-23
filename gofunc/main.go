package main

import (
	"fmt"
)

func main()  {
	// 使用协程执行run函数（异步执行）
	go run()

	fmt.Println(222)
}

func run()  {
	fmt.Println(111)
}