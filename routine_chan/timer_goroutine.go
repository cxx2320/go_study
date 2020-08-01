package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1e8)  // 每1秒向tick写入时间
	boom := time.After(5e8) // 5秒后向boom写入时间

	// 无限循环执行select
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			// 这里进行了return 当走到这里时会结束for循环
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
