package main

import (
	"fmt"
	"math"
)

/*
Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
*/

func computeWithParameters(fn func(int, int) int, x, y int) int {
	return fn(x, y)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//adder return a function which get a int value as a parameter and return int value
func adder() func(int) int {
	sum := 0
	fmt.Println("Created adder")
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println("With parameter")
	divider := func(x, y int) int {
		return x / y
	}

	fmt.Println(computeWithParameters(divider, 30, 6))

	fmt.Println()
	fmt.Println("Adder start...")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
