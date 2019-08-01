package _interface

import (
	"fmt"
	"math"
)

type In interface {
	M()
}

type Truck struct {
	S string
}

func (t *Truck) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i In

	i = &Truck{"Hello"}
	describe(i)
	i.M()
	fmt.Println()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i In) {
	fmt.Printf("(%v, %TStruct)\n", i, i)

}
