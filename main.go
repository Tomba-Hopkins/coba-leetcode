package main

import (
	"fmt"
	"math"
)

func squirt(num int) bool {
    sqr := math.Sqrt(float64(num))
	
	exp := int(sqr)

	return exp * exp == num
}

func main() {
	fmt.Println(squirt(81))
}