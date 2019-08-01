package main

import "fmt"

var power = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range power {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	fmt.Println("")
	bok := make([]int, 5)
	for i := range bok {
		bok[i] = 1 << uint(i) // == 2**i
	}

	for _, value := range bok {
		fmt.Printf("%d\n", value)
	}
}
