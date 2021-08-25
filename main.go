package main

import (
	"fmt"
	"math"

	ar "sapan.com/arrays"
	"sapan.com/maps"
)

// "sapan.com/controlstructure"
// "sapan.com/maps"
// "sapan.com/slice"
// "sapan.com/variables"

type circle struct {
	radius float64
}

type rect struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r rect) area() float64 {
	return r.length * r.width
}

func main() {
	// variables.Variables()
	// controlstructure.Controlstructure()
	// arrays.Arrays()
	// slice.Slice1()
	// maps.Tst_maps()
	// var1 := 1
	// pointers.WithoutReference(var1)
	// fmt.Println(var1)
	// var2 := 2
	// pointers.WithReference(&var2)
	// fmt.Println(var2)
	c1 := circle{2}
	r1 := rect{2, 4}
	shapes := []shape{c1, r1}
	for _, v := range shapes {
		fmt.Println(v.area())
	}
	v1 := ar.S1{maps.S2{1, 2}, 5.6}
	fmt.Println(v1)

}
