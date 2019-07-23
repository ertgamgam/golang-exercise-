package main

import "fmt"

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
}
