package main

import "fmt"

func main() {
	manyMap()
}

func makeMap() {
	var mapLit = map[string]string{}
	fmt.Println(mapLit)
}

func makeMap1() {
	// 使用 interface 声明一个值可以为任意类型的map
	// mapList := map[string]interface{"name":"cxx","age":18}

	//初始化
	// var mapLit = map[string]string

	// 声明一个空map
	user := make(map[string]string)

	// 声明一个空map
	// user := map[string]string{}

	// 声明一个有值的map
	// user := map[string]string{"age": "18"}

	user["name"] = "柴星星"
	user["age"] = "18"

	fmt.Println("年龄", user["age"])

	capital, ok := user["age"]
	fmt.Println(capital, ok)

	// range循环 map
	for key, v := range user {
		fmt.Println(key, v)
	}

	// 删除
	delete(user, "age")
	age, ok := user["age"]
	fmt.Println(age, ok)
}

// 检查key是否存在
func checkKeyExits() {
	mapList := map[string]int{"name": 20, "age": 18}
	if _, ok := mapList["name"]; ok {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}
}

// 删除key
func deleteKey() {
	mapList := map[string]int{"name": 20, "age": 18}
	delete(mapList, "name")
	if _, ok := mapList["name"]; ok {
		fmt.Println("exists")
	} else {
		fmt.Println("not exists")
	}
}

func q8() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}

// 多值map
func manyMap() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("Version A: Value of items: %v\n", items)
}
