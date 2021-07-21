package pointers

import "fmt"

func WithoutReference(x int) int {
	fmt.Println(x)
	x = 90
	return x
}

func WithReference(x *int) int {
	fmt.Println(*x)
	*x = 90
	return *x
}
