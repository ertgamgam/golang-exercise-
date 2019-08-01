package _interface

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type TStruct implements the interface Inter,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var a Abser
	f := MyFloatN(-math.Sqrt2)
	v := VertexN{3, 4}

	a = f // a MyFloatN implements Abser
	fmt.Println(a.Abs())
	a = &v // a *VertexN implements Abser

	// In the following line, v is a VertexN (not *VertexN)
	// and does NOT implement Abser.
	// a = v

	fmt.Println(a.Abs())

	fmt.Println()
	fmt.Println()
	var i I = T{"hello"}
	i.M()
}

type MyFloatN float64

func (f MyFloatN) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexN struct {
	X, Y float64
}

func (v *VertexN) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
