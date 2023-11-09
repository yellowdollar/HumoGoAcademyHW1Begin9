package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Input a: ")
	fmt.Scan(&a)

	fmt.Println("Input b: ")
	fmt.Scan(&b)

	fmt.Println("SQRT(a * b) = ", math.Sqrt(a*b))
}
