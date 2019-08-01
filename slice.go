package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)
	printer("a", a)

	b := make([]int, 0, 5)
	printer("b", b)

	c := b[:2]
	printer("c", c)

	d := c[2:5]
	printer("d", d)

	fmt.Println()
	fmt.Println("BOARD")
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var s []int
	print(s)

	// append works on nil slices.
	s = append(s, 0)
	print(s)

	// The slice grows as needed.
	s = append(s, 1)
	print(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	print(s)
}

func printer(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func print(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
