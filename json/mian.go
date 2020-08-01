package main

import (
	"encoding/json"
	"fmt"
)

// User 模型
type User struct {
	Name  string
	Age   int
	Price float64
}

func main() {
	JosnTo()
}

func ToJosn() {
	user := User{"Cxx", 18, 8888}
	fmt.Printf("%s \n", user)

	// 结构体转json
	json1, _ := json.Marshal(user)
	// 需要进行格式化输出才能看到效果
	fmt.Printf("%s \n", json1)
}

func JosnTo() {
	b := []byte(`{"Name": "Cxx", "Age": 6, "Price": 2584.6}`)
	var f User
	json.Unmarshal(b, &f)
	fmt.Println(f)
}
