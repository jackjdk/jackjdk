package main

import "fmt"

func fun_three() {
	fmt.Println("俺是一个fun_three函数")
}

func main() {
	/*
		匿名：没有名字
			匿名函数：没有名字的函数
		定义一个匿名函数，直接进行调用。通常只能使用一次.也可以使用匿名函数赋值给某个函数变量，那么就可以使用多次了

		匿名函数：
			Go语言是支持函数式编程：
			1.将匿名函数作为另一个函数的参数，回调函数
			2.将匿名函数作为另一个函数的返回值，可以形成闭包结构
	*/

	fun_two := fun_three
	fun_three()
	fun_two()

	//匿名函数
	func() {
		fmt.Println("俺是一个匿名 函数")
	}()

	fun3 := func() {
		fmt.Println("俺又是一个匿名函数呢")
	}
	fun2 := fun3 //将fun3的值赋给fun2
	fun2()
	fun3()

	//var myFunc1 func()
	//myFunc1 = func() {
	//	fmt.Println("Hello, World!")
	//}
	//myFunc1()
	//var myFunc2 func()
	//myFunc2 = func() {
	//	fmt.Println("hello")
	//}
	//myFunc2()

	/*
		//定义带参数的匿名函数
		func(a, b int) {
			fmt.Println(a, b)
		}(1, 2)

		//定义带返回值的匿名函数
		res1 := func(a, b int) int { //在匿名函数结尾的括号里传入值给return，返回给res1 ，打印
			return a + b
		}(10, 20)
		fmt.Println(res1)

		res2 := func(a, b int) int {
			return a + b
		} //将匿名函数的值(地址)，赋值给res2
		fmt.Println(res2)

		fmt.Println(res2(100, 200))

	*/
}
