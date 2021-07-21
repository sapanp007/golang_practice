package slice

import (
	"fmt"
	"sort"
)

func Slice1() {
	s1 := make([]int, 0)
	var s2 []int
	fmt.Println(s1, s2)

	//sorting
	s3 := []int{99, 2, 3, 4, 5}
	sort.Ints(s3)
	fmt.Println(s3)
	// sort.Float64s
	// sort.Strings

	family := []struct {
		Name   string
		Age    int
		chocos int
	}{
		{"Alice", 23, 7},
		{"David", 2, 9},
		{"Eve", 2, 9},
		{"Bob", 25, 1},
	}
	// Sort by age, keeping original order or equal elements.
	sort.Slice(family, func(i, j int) bool {
		return family[i].chocos > family[j].chocos
	})
	fmt.Println(family)
}
