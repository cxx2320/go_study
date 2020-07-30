package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	// 实现一个接口的方式一
	// var areaIntf Shaper
	// areaIntf = sq1

	// 实现一个接口的方式二
	areaIntf := Shaper(sq1)

	// ( 这种继承方式报错
	// or even:
	// areaIntf := sq1
	// )
	fmt.Printf("The square has area: %f\n", areaIntf.Area())
}
