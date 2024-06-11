package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := [] int {4,6,7,9,8}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"test","123","xxx","kkk","zzz"}
	sort.Strings(names)
	fmt.Println(names)


	heights := [] float64{1.1,-1.1,3.3,2.2}
	sort.Float64s(heights)
	fmt.Println(heights)
}
