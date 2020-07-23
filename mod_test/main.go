package main

/**
 * go mod 使用
 */

import (
	"fmt"
	"mod_test/test_math"

	// 使用别名
	String "mod_test/string"

	"github.com/jinzhu/configor"
)

func main() {
	fmt.Println("test", configor.Config{})
	fmt.Println("test1", test_math.Add(200, 30))
	fmt.Println("test2", String.Add("Cxx"))
}
