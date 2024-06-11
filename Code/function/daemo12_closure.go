package main

import "fmt"

/*
Go语言支持函数式编程

	支持将一个函数作为另一个函数的参数，
	也支持将一个函数作为另一个函数的返回值
	闭包(closure):
		一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或者外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数
		这个内层函数和外层函数的局部变量，统称为闭包结构

		局部变量的声明周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁
		但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用
*/
func ADDpper() func(int) int { //外层函数。声明一个不接受任何参数的函数。并接受一个int类型的参数和返回一个int类型的结果
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func increment() func() int { //外层函数
	i := 0              //局部变量
	fun := func() int { //内层函数。定义了一个匿名函数，给变量自增并返回
		i++
		return i
	}
	return fun
}

func main() {

	bibao := increment() //bibao= fun
	fmt.Printf("%T\n", bibao)
	fmt.Println(bibao)
	//b1 := bibao()
	//fmt.Println(b1)
	//b2 := bibao()
	//fmt.Println(b2)
	//
	//f := ADDpper()
	//fmt.Println(f(3))
	//add2 := func(n int) int {
	//	return n + 2
	//}
	//add10 := func(n int) int {
	//	return n + 10
	//}
	//add100 := func(n int) int {
	//	return n + 100
	//}

	addBase := func(base int) func(int) int {
		return func(n int) int {
			return base + n
		}
	}
	//fmt.Println(add2(3))
	//fmt.Println(add10(3))
	//fmt.Println(add100(3))

	add8 := addBase(8)
	fmt.Println(add8(2))
}
