package main

import "fmt"

// https://learnku.com/docs/the-way-to-go/72-slice/3613

func main() {
	// 7.3
	s := make([]byte, 5)
	// fmt.Println(cap(s), len(s)) // 5 5

	s = s[2:4]
	// fmt.Println(cap(s), len(s)) // 3 2

	// 7.4
	s1 := []string{"p", "o", "e", "m"}
	s2 := s1[2:]
	// fmt.Println(s2) // [e m]

	s2[1] = "t"
	fmt.Println(s1) // [p o e t]
	fmt.Println(s2) // [e t]
}
