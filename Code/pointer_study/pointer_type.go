package main

import "fmt"

func main() {
	/*数组指针:
	首先是一个数组，一个数组的地址。
	*[4]Tyep

	指针数组：首先是一个数组，存储的数据类型是指针
	[4]*Type

	*[5]flout64, 指针，一个存储了5个浮点类型数据的数组的指针
	*[3]string, 指针，数组的指针，存储了3个字符串
	[3]*string, 数组，存储了3个字符串的指针地址的数组
	[5]*flout64, 数组，存储了5个浮点数的地址的数组
	*[5]*flout64, 指针，存储了5个浮点类型的数据的指针地址的数组的指针
	*[3]*string，指针，存储了3个字符串的指针地址的数组的指针
	**[4]string, 指针，存储了4个字符串数据的数组的指针的指针
	**[4]*string，指针，存储了4个字符串指针地址的数组，的指针的指针*/
	/*
		//1.创建一个普通的数组
		arr1 := [4]int{1, 2, 3, 4}
		fmt.Println(arr1)

		//2.创建一个指针，存储该数组的地址-->数组指针
		var p1 *[4]int
		p1 = &arr1
		fmt.Println(p1)         //&[1 2 3 4]
		fmt.Printf("%p\n", p1)  //数组arr1的地址
		fmt.Printf("%p\n", &p1) //p1指针自己的地址

		//3.根据数组指针，操作数组索引的元素并修改为100
		(*p1)[0] = 100
		fmt.Println(arr1)

		p1[0] = 200 //简化写法
		fmt.Println(arr1)
	*/
	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}
	fmt.Println("输出arr2的值", &arr2)
	//根据arr2数组的索引获取地址
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println(&arr3)
	//arr4 := &arr3
	var arr4 **int

	//arr4 := &arr3
	fmt.Printf("%p\n", &arr4)
	fmt.Println("输出arr4", arr4)

	fmt.Println()
	fmt.Println(arr2) //[1 2 3 4]
	fmt.Println(arr3) //arr3取的地址
	//fmt.Println(*arr3[0])
	arr2[0] = 100
	fmt.Println(arr2)
	fmt.Println(a)
	/*
		*arr3[0] = 200
		fmt.Println(arr3)
		fmt.Println(a)



		b = 1000
		fmt.Println(arr2)
		fmt.Println(arr3)

	*/

}
