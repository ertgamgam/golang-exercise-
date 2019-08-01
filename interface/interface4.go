package _interface

import "fmt"

func main() {
	var i interface{}
	describe3(i)

	i = 42
	describe3(i)

	i = "hello"
	describe3(i)
}

func describe3(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
