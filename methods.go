package main

import (
	"fmt"
	"math"
)

type VertexM struct {
	X, Y float64
}

func (v VertexM) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*With a value receiver, the Scale method operates on a copy of the original Vertex value.
	(This is the same behavior as for any other function argument.)
The Scale method must have a pointer receiver to change the Vertex value declared in the main function.*/
func (v *VertexM) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *VertexM, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//can declare a method on non-struct types, too.
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := VertexM{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	fmt.Println()
	fmt.Println()
	fmt.Println()

	ve := VertexM{3, 4}
	ve.Scale(2)
	ScaleFunc(&ve, 10)

	p := &VertexM{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(ve)
	fmt.Println(p)


}

/*There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)*/
