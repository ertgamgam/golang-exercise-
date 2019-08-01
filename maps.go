package main

import "fmt"

type VertexMap struct {
	Lat, Long float64
}

func main() {
	var m map[string]VertexMap
	m = make(map[string]VertexMap)
	m["Bell Labs"] = VertexMap{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var mp = map[string]VertexMap{
		"Bell Labs": {40.68433, -74.39967,},
		"Google":    {37.42202, -122.08408,},
	}

	fmt.Println(mp["Google"])
	fmt.Println(mp)

	fmt.Println("insert")
	mp["Facebook"] = VertexMap{Lat:12.33,Long:23.44}
	fmt.Println(mp["Facebook"])

	fmt.Println("mapper")
	mapper :=make(map[string]int)
	mapper["age"] = 33
	fmt.Println("before update = ",mapper["age"])
	mapper["age"]=34
	fmt.Println("after update = ",mapper["age"])

	v,ok := mapper["age"]
	fmt.Println("Check before deleting... The value:", v, "Present?", ok)

	delete(mapper,"age")
	vAfterDelete,okAfterDelete := mapper["age"]
	fmt.Println("Check after deleting... The value:", vAfterDelete, "Present?", okAfterDelete)

}
