package main

import (
	"errors"
	"fmt"
)

func main() {
	test()
}

func test() {
	defer CoverErrorMessage()
	panic("panicasd")
	fmt.Println("success")

	fmt.Println("successsss")
}

func CoverErrorMessage() {
	if message := recover(); message != nil {
		var err error
		switch x := message.(type) {
		case string:
			err = errors.New(x)
		case error:
			err = x
		default:
			err = errors.New("Unknow panic")
		}
		fmt.Println(err)
	}
}
