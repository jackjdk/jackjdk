package main

import "fmt"

func main() {
	var slice []int
	slice = append(slice, 8)
	slice = append(slice, 7)
	fmt.Println(slice)
	sc := slice[0:2] //通过将
	fmt.Println(sc[0], sc[1])

	//sc = 10
	//
	//fmt.Println(sc)
	//slice[0] = sc //将sc的值赋给slice的指定元素
	//
	//fmt.Println(slice)
}
