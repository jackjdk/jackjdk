package main

import "fmt"

func main() {
	/*

		作用域：变量可以使用的范围。一段代码的{}就是作用域
			局部变量：函数内部定义的变量，就叫做局部变量
			全局变量，函数外部定义的变量，就叫做全局变量
	*/
	//在定义main函数中，所以n的作用域就是main函数的范围内

	n := 10

	if a := 1; a <= 10 {
		fmt.Println(a)
		fmt.Println(n)
	}
	//	fmt.Println(a) //不能访问a，出了作用域
	fmt.Println(n)

	if b := 1; b <= 40 {
		n := 20
		fmt.Println(n)
		fmt.Println(b)
	}

}
