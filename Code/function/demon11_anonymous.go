package main

import "fmt"

//函数声明 定义,add=标识符
//()=定义时，声明一个函数。未来某个时刻，需要使用的时候。也可以不用。可以用N次

// func add(x int,y int)  {//函数签名，不完整。
func add(x int, y int) int { //带返回值签名。xy形参代表未来必须在调用时提供2个对应的实参，int表示返回值的个数和类型。返回值类型必须确定
	r := x + y
	fmt.Println(r) //xy 称为形式上的参数，简称形参。Go语言中有些人称为入参。形参是占位用的
	return r       //返回值。1个返回值，类型是int.不带return默认返回的是零值
}

//func add() {
//	fmt.Println("函数")
//}

//func addN(a,b int, args ...int) int {
//	return 0
//}

// 格式化
func print(callback func(...string), args ...string) {
	fmt.Println("print函数输出:")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, ":", v)
	}
}

func main() {
	//add() //调用时:调用called指定是使用函数标识符或函数本身，在后面加上小括号
	//add(4,5)调用了一次called一次。4,5一次和xy对应，这是在未来的某个时刻我实实在在提供的具体数据。调用时需要提供实际的参数，这个参数称为实际参数，简称实参
	//var b func(int,int)int = add //也可以通过var变量来设置值
	d := add(4, 5) //调用时，如果有返回值，使用对应个数变量接收。用4和5来代替形参的过程，叫做传参。
	//fmt.Println(b(10,40))
	fmt.Println(d)
	fmt.Println(add(10, 20) + 3 + 6)

	//const v1 = add(1,2) //不行。Go语言不支持常量调用函数
	//var v1 = add(1,2)
	//fmt.Println(v1 + 100)

	//print(list."A","C","E")

	print(list, "A", "C", "E")

	////匿名函数
	//匿名：没有名字
	//匿名函数：没有名字的函数
	//定义一个匿名函数，直接进行调用。通常只能使用一次.也可以使用匿名函数赋值给某个函数变量，那么就可以使用多次了
	//
	//匿名函数：
	//Go语言是支持函数式编程：
	//1.将匿名函数作为另一个函数的参数，回调函数
	//2.将匿名函数作为另一个函数的返回值，可以形成闭包结构

	sayHello := func(name string) {
		fmt.Println("Hello", name)
	}
	sayHello("KK")

	func(name string) { //定义了一个匿名函数，它接受一个类型为 string 的参数 name，并输出 "Hi" 和该参数。
		fmt.Println("Hi", name)
	}("KK")

	values := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}
	print(values, "A", "B", "C")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Println(v, "\t")
		}
		fmt.Println()
	}, "A", "C", "E")
}
