package main

import "fmt"

func main() {
	/*
		高阶函数：
			根据Go语言的数据类型的特点，可以将一个函数作为另一个函数的参数
		fun1(),fun2()
		将fun1函数作为了fun2这个函数的参数。

				fun2函数: 就叫高阶函数
					接收了一个函数作为参数的函数，高阶函数
				fun1函数：回调函数
					作为另一个函数的参数的函数，叫做回调函数
	*/
	//加法运算
	fmt.Printf("%T\n", Add)  //func(int, int) int
	fmt.Printf("%T\n", oper) //func(int, int, func(int, int) int) int

	res1 := Add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 40, Add)
	fmt.Println(res2)

	fun_one := func(a, b int) int {
		return a * b
	}

	res4 := oper(20, 10, fun_one)
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			return 0
		}
		return a / b
	})
	fmt.Println(res5)

}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}
func Add(a, b int) int {
	return a + b
}
