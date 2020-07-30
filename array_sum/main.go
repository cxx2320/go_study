package main

import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x, cxx := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Println("The sum of the array is: %f", x)
	fmt.Println(cxx)
}

// 如果返回值有变量，无需放在return后面就能返回
func Sum(a *[3]float64) (num float64, cxx string) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		num += v
	}
	cxx = "Dxx"
	return
}
