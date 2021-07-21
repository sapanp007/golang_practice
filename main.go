package main

import (
	"fmt"
	// "sapan.com/arrays"
	// "sapan.com/controlstructure"
	// "sapan.com/maps"
	// "sapan.com/slice"
	// "sapan.com/variables"
	"sapan.com/pointers"
)

func main() {
	// variables.Variables()
	// controlstructure.Controlstructure()
	// arrays.Arrays()
	// slice.Slice1()
	// maps.Tst_maps()
	var1 := 1
	pointers.WithoutReference(var1)
	fmt.Println(var1)
	var2 := 2
	pointers.WithReference(&var2)
	fmt.Println(var2)
}
