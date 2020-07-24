package main

import "fmt"

func main() {

	// 声明一个空map
	user := make(map[string]string)

	// 声明一个空map
	// user := map[string]string{}

	// 声明一个有值的map
	// user := map[string]string{"age": "18"}
	user["name"] = "柴星星"
	user["age"] = "18"

	fmt.Println("是", user["age"])

	capital, ok := user["age"]
	fmt.Println(capital, ok)

	// range循环 map
	for key := range user {
		fmt.Println("key:", key)
	}

	// 删除
	delete(user, "age")
	age, ok := user["age"]
	fmt.Println(age, ok)
}
