package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	// 返回复制成功的条数
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl4 := []int{7, 55, 99, 3}
	// 返回追加后的切片(后一个切片后面需要加上三个点)
	sl3 = append(sl3, sl4...)
	fmt.Println(sl3)
}
