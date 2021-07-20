package main

import "fmt"

func main() {
	var x int = 89

	// if-else block
	if x == 90 {
		fmt.Println("x is 90")
	} else if x == 78 {
		fmt.Println("x is 78")
	} else {
		fmt.Println("x is 89")
	}

	//Switch Block
	switch x {
	case 78:
		fmt.Println("x is 78")
	case 90:
		fmt.Println("x is 90")
	default:
		fmt.Println("x is 89")
	}
}
