package main

import (
	"errors"
	"fmt"
)

var errorNot error = errors.New("Not found error")

func main() {
	// fmt.Printf("Error: %s\n", errors.New("not found"))
	fmt.Errorf("math: square root of negative number")
}
