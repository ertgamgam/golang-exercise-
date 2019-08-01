package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	s := primes[1:5] //or 	var s []int = primes[1:5]
	fmt.Println(s)

	fmt.Println()
	fmt.Println()

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	k := names[0:2]
	l := names[1:3]
	fmt.Println(k, l)

	l[0] = "XXX"
	fmt.Println(k, l)
	fmt.Println(names)

	fmt.Println()
	fmt.Println()
	//
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	m := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(m)

	newValue := struct {
		i int
		b bool
	}{
		19, true,
	}
	fmt.Println("After added new value to struct")
	m = append(m, newValue)
	fmt.Println(m)

	//
	fmt.Println("slice defaults")
	slice := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice)

	slice = slice[1:4]
	fmt.Println(slice)

	slice = slice[:2]
	fmt.Println(slice)

	slice = slice[1:]
	fmt.Println(slice)

	fmt.Println("The zero value of a slice is nil." +
		"A nil slice has a length and capacity of 0 and has no underlying array.")

	var emptySlice []int
	fmt.Println(emptySlice, len(emptySlice), cap(emptySlice))
	if emptySlice == nil {
		fmt.Println("nil!")
	}

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
