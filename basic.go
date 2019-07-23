package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

//A function can return any number of results.
func swap(x, y string) (string, string) {
	return y, x
}

//Go's return values may be named. If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

var k, l = 1, 2 //or var k, l int = 1, 2

func main() {
	fmt.Println(add(3, 5))

	fmt.Println(add2(3, 5))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(18))

	var i int
	fmt.Println(i, c, python, java)

	check, control, java := true, false, "yesorNO!" //or  var check, control, java = true, false, "no!"
	fmt.Println(k, l, check, control, java)

	//Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	//Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
}
