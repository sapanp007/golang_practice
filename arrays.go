package main

import "fmt"

func main() {
	fmt.Println("Learn more about arrays in Go")
	// var a1 [5]int
	// a1 = [5]int{1, 2, 3, 4, 5}

	a1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(a1), cap(a1))

	a2 := a1[1:]
	fmt.Println(len(a2), cap(a2))

	a3 := a1[2:4]
	fmt.Println(len(a3), cap(a3))
	//a3[3] = 90	// Index out of range
	fmt.Println(a3)

	x := make([]int, 0, 3)
	fmt.Println(x)
	x = append(x, 1)
	fmt.Println(x)
	x = append(x, a3...) //... is essential
	fmt.Println(x)

	slice1 := make([]int, 2)
	slice2 := make([]int, 1, 1)
	fmt.Println(slice2, len(slice2), cap(slice2))
	slice2 = append(slice2, 1, 3)
	fmt.Println(slice2, len(slice2), cap(slice2))

	//copy slice1 into slice2
	total_elem_copied := copy(slice2, slice1)
	//total_elem_copied will give int value of Total number of elements copied
	//Note: items always gets copied from index 0
	fmt.Println(total_elem_copied, slice2, slice1)
}
