package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {

	//for loop
	adder := 0
	for i := 0; i < 10; i++ {
		adder += i
	}
	fmt.Println(adder)

	sum := 1
	for sum < 1000 { //or for ;sum < 1000;
		sum += sum
	}
	fmt.Println(sum)

	/* infinite loop
	for {
	}*/

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(runtime.GOOS)

	//defer

	defer fmt.Println("world")
	fmt.Println("hello")

	defer fmt.Println("GAMGAM")
	defer fmt.Println("Ertuğrul")
	fmt.Println("test")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
