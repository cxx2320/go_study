package main

import "fmt"

func main() {
	items := [...]int{10, 20, 30, 40, 50}

	// 在循环内修改value不会作用到原始变量items
	for i, item := range items {

		// 这样可以修改items的值
		items[i] *= 2

		// 这样不行
		// item *= 2
	}
	fmt.Println(items)
}