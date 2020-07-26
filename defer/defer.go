package main

import "fmt"

// defer调用顺序，先声明的后调用（栈，先进后出）。init函数也采用了这样的方法

func main() {
    fmt.Printf("main")
    defer aa()
    bb()
    defer bb()
}

func aa(){
    fmt.Printf("A")
}

func bb(){
    fmt.Printf("B")
}