package _interface

import "fmt"

type Inter interface {
	M()
}

type TStruct struct {
	S string
}

func (t *TStruct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i Inter

	var t *TStruct
	i = t
	describe2(i)
	i.M()

	i = &TStruct{"hello"}
	describe2(i)
	i.M()
}

func describe2(i Inter) {
	fmt.Printf("(%v, %TStruct)\n", i, i)
}
