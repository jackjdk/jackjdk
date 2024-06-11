package main

import "fmt"

func main() {
	var a int
	var b byte
	var f float32
	var t bool
	var s string
	var r rune
	var arr []int

	fmt.Println("default value of  int %d\n", a)
	fmt.Println("default value of  byte %d\n", b)
	fmt.Println("default value of  float32 %f\n", f)
	fmt.Println("default value of  bool %t\n", t)
	fmt.Println("default value of  strint [%s]\n", s)
	fmt.Println("default value of  int %d %c\n", r, r)
	fmt.Println("default value of slice %v\n", arr)

}
