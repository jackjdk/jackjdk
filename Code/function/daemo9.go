package main

import "fmt"

func main() {
	/*
		Go语言的数据类型
			基本数据类型：
					int、float、bool、string
			符合数据类型：
					array、slice、map、function、pointer、struct、interface
			数值类型：整数、浮点
				进行运算操作、加减乘除、打印
			字符串：
				可以获取单个字符串，截取子串，遍历，string包下的函数操作。
			数组、切片、map。
				存储数据、修改数据，获取数据，遍历数据
			函数：加(),进行调用


			注意点：
				函数作为一种复合数据类型，可以看做是一种特殊的变量
				函数名():将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
				函数名：指向函数体的内存地址
	*/
	//1.整形
	a := 10
	//运算
	a += 15

	a += 20

	fmt.Println("a:", a)

	//2.数组、切片、map.容器
	b := []int{1, 2, 3, 4}
	b[0] = 100
	fmt.Println(len(b))
	fmt.Println("b:", b)
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d\t", b[i]) //通过循环遍历i变量的元素，前提i的值小于len(长度)
	}
	fmt.Println()

	//3.函数做一个变量
	fmt.Printf("%T\n", fun_one) //func(int,int)

	fmt.Println(fun_one) //fun_ond  0xc8fa80,看做函数名对应的函数体的地址

	//4.直接定义一个函数类型的变量
	var c func(int, int)
	fmt.Println(c)

	//var d string
	//d = "hello"
	c = fun_one
	fmt.Println(c)
	fun_one(10, 20)
	c(100, 300) //c也是函数类型，加小括号也可以使用

	res1 := fun_two       //将fun_two的值(函数的地址)赋值给res1，res1和fun2指向同一个函数体
	res2 := fun_two(1, 2) //将fun_two函数进行调用，将函数的执行结果赋值给res2.相当于：a+b

	fmt.Println(res1)
	fmt.Println(res2)

	fmt.Println(res1(10, 30)) //也可以被调用

}

func fun_two(a, b int) int {
	return a + b
}

func fun_one(a, b int) {
	fmt.Println("a:%d,b:%d\n", a, b)

}
