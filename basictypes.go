package main

import (
	"fmt"
	"math/cmplx"
)

/*
<<   left shift             integer << unsigned integer
>>   right shift            integer >> unsigned integer*/

// (x << n == x*2^n ) (x >> n == x*2^(-n))

//The expression T(v) converts the value v to the type T.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const Pi = 3.14 //or const Pi float32= 3.14

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
