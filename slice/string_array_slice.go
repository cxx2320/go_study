package main

import "fmt"

func main() {

	s := "hello"
	fmt.Println(s)

	c := []byte(s)
	fmt.Println(c)

	c[0] = 'c'
	fmt.Println(c)

	s2 := string(c)
	fmt.Println(s2) // s2 == "cello"

	// 可以通过单引号打印出0~9，a~Z 的ASCII码
	fmt.Println('A')
}
