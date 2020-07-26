
package main

import "fmt"

func main() {

	// p2 为Add2函数执行后的返回值
	/*
	是一个函数
	func(b int) int {
        return b + 2
	}
	然后使用p2函数传入参数3，最终结果即为5
	*/
    p2 := Add2()
    fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	
	// 和上面类似，只不过最外层的函数接收一个a变量，暂时的存储了起来，
	// 在调用Adder（2）函数时返回 TwoAdder 函数，此时调用此函数外部的a变量可以被内层函数访问到
    TwoAdder := Adder(2)
    fmt.Printf("The result is: %v\n", TwoAdder(3))
}

func Add2() func(b int) int {
    return func(b int) int {
        return b + 2
    }
}

// 如果在嵌套一层函数会怎么样?
func Adder(a int) func(b int) int {
    return func(b int) int {
        return a + b
    }
}
