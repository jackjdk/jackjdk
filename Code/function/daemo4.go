package main

import "fmt"

//全局变量的定义
//num3 := 1000 //不支持简短定义的写法
const num3 = 1000
func main()  {
	/*

		作用域：变量可以使用的范围
			局部变量：函数内部定义的变量，就叫做局部变量
					变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了
			全局变量：函数外部定义的变量，就叫做全局变量
					所有的函数都可以使用，而且共享这一份数据
	*/
	//在定义main函数中，所以n的作用域就是main函数的范围内

	n := 10

	if a := 1;a <= 10 {
		fmt.Println(a)
		fmt.Println(n)
	}
	//	fmt.Println(a) //不能访问a，出了作用域
	fmt.Println(n)

	if b := 1; b<=40{
		n:= 20
		fmt.Println(n)
		fmt.Println(b)
	}
	fun()
	fun3()

fmt.Println("main中访问全局变量",num3)
}


func fun(){
	num1 := 100
	fmt.Println("fun()函数中:num1:",num1)
	//num3 = 1000			//常量定义了全局变量定义了之后只能被调用，不能被修改常量的值，所以会报错
	fmt.Println("fun()函数,访问全局变量 num1:",num3)
}

var num1  = 100
func fun3(){
	num1 := 200
	fmt.Println("fun()函数中:num1:",num1)
	//num3 = 1000
	fmt.Println("func3函数，访问全局变量 num1:",num3)
}






