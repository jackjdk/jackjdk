package main

import (
	"fmt"
)

func main() {
	/*
		指针：pointer
		存储了另一个变量的内层地址的变量。
		& 取要打印某个值的内存地址
		* 表示指向一个指针的整数、字符串、函数、切片等等类型
		*是创建指针。然后再使用*去引用创建的指针。&是用来给创建的指针引用变量中的内存地址，根据内存地址使用*来获取值
		A :=2
		B :=A
		B =3
		fmt.Println(A,B)


		//指针
		var cc *int	//创建一个空指针
		cc = &A    //指针c取的A的内存地址
		C := &A
		fmt.Printf("%T %T\n",C,cc)
		fmt.Println(*cc)  //这时候指针打印的还是A的值

		*cc = 4  //现在要给cc修改值
		fmt.Println(A)
	*/
	v1 := 19

	//创建一个指针变量，用于存储变量a的地址
	var v2 *int
	//<nil> 空指针
	fmt.Println(v2)
	//获取v1的内存地址，并将其赋值给v2指针
	v2 = &v1
	fmt.Println(*v2)
	// 使用*操作符修改v1的值
	*v2 = 29
	// 再次访问v1的值，并打印出来
	fmt.Println(*v2)
	fmt.Printf("%T\n", *v2, &v2)
	fmt.Println(&v1)

	v1 = 100
	fmt.Println(v1)
	//指向指针的指针
	var p3 **int
	fmt.Println(p3)
	p3 = &v2
	fmt.Println(p3, v1, v2)
	fmt.Printf("%p\n", &p3, &v2, &v1)
	fmt.Println(*p3)

	/*
		fmt.Println(v2)
		fmt.Printf("%p %T\n", v1, v2)
		fmt.Println("v2的数值是:", v2)                 //v2存储的是v1的地址
		fmt.Printf("v2自己的内存地址:%p\n", &v2)          //打印自己的内存地址
		fmt.Println("v2的数值，是v1的地址，该地址存储的数据:", *v2) //打印自己内存地址的值
		/*
		   	//操作变量，更改数值,并不会改变地址
		   	v1 = "ABC2"
		   	fmt.Println(v1)
		   	fmt.Printf("%p\n", &v1)

		   	//通过指针，改变变量的数值
		   	*v2 = "bce3"
		   	fmt.Println(v1)

		   	//指针的指针
		   	var v3 **string
		   	fmt.Println(v3)
		   	v3 = &v2
		   	fmt.Printf("%T，%T，%T\n", v1, v2, v3)
		   	fmt.Println("v2的数值:", v2)
		   	fmt.Printf("v2自己的地址: %p\n", &v2)

		   	//v1 = "alex"
		   	//fmt.Println(v1,v2,*v2)

	*/

}
