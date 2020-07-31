package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	fmt.Println("Please enter some input: ")
	inputReader = bufio.NewReader(os.Stdin)

	// 当遇到s后，后面的不再进行监听
	// 如在键盘键入 aaasss
	// input的值为 aaas
	input, err = inputReader.ReadString('s')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
