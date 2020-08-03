package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timeout := make(chan bool, 1)

	// 一直在ch里面写值
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	// 一秒后在timeout里面写入值
	go func() {
		time.Sleep(1e9 * 1) // one second
		timeout <- true
	}()

	// 无限在ch和timeout读取值，当读取到timeout的值时退出
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
