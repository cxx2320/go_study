package main

import (
	"fmt"
)

type tryFunc func()
type catchFunc func(interface{})

func Try(try tryFunc, catch catchFunc) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()
	try()
}

func main() {
	Try(func() {
		panic("1")
		panic("2")
	}, func(a interface{}) {
		fmt.Println("ww", a)
	})
}
