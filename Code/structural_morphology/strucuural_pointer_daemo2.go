package main

import (
	"fmt"
)

type Person2 struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {
	/*
				数据类型:
					值类型: int	float	bool	string	array

					引用类型: slice	map	 function	pointer

				通过指针：
					new()，不是nil，是一个初始化后的空指针
		new()是一个内置函数，用于分配内存并返回指向该内存的指针。new()函数只分配内存，并将其初始化为零值，但不会对分配的内存进行进一步的初始化。
		这意味着，如果你使用new()创建了一个类型的实例，那么你将得到一个指向该类型的新实例的指针，但该实例的字段将被初始化为零值。

		与new()不同，nil是一个特殊的预定义标识符，表示一个空指针或空接口值。在Go语言中，nil用于表示指针、切片、映射、通道和函数类型的零值或空值。
		当你声明一个指针类型的变量时，它的初始值就是nil，表示该指针不指向任何有效的内存地址。

			nil表示指针的零值或空值，它不指向任何有效的内存地址。
			new()用于分配内存并返回一个指向该内存的指针，但不会对内存进行进一步的初始化。
	*/
	//1.结构体是值类型
	p1 := Person2{name: "王二狗", age: 30, sex: "男", address: "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p %T\n", &p1, p1)

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	p2.name = "李小白"
	fmt.Println("打印p2通过p1取的Person2值", p2)
	fmt.Println(p1)

	//2.定义结构体指针指向了最开始的结构体
	var pp1 *Person2
	pp1 = &p1
	fmt.Println("打印指针取的值", pp1)
	fmt.Printf("打印指针地址和类型%p,%T\n", pp1, pp1)
	fmt.Println(*pp1)

	//(*pp1).name = "李四"通过*
	pp1.name = "王五"
	fmt.Println(pp1)
	fmt.Println(p1)

	//使用内置函数new()，go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person2)
	fmt.Printf("打印new%T\n", pp2)
	fmt.Println(pp2)
	//(*pp2).name
	pp2.name = "jerry"
	pp2.age = 20
	pp2.sex = "男"
	pp2.address = "上海市"
	fmt.Println(pp2)

}
