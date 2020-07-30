package main

import "fmt"

type IntArray []int

func (this IntArray) Len() int {
	return len(this)
}

// 这里为什么会改变自身的值
func (this IntArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	Arr := IntArray{1, 5, 6, 9, 8}
	Arr.Swap(0, 1)
	fmt.Println(Arr)
}
