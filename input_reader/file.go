package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, _ := ioutil.ReadFile("./test.txt")
	fmt.Println(string(content))
	ioutil.WriteFile("./test.txt", []byte("Ni HAO"), 0644)
}
