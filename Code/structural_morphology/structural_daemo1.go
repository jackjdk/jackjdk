package main

import "fmt"

func main() {
	/*
		结构体:结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。
			结构体成员是有一系列的成员变量构成，这些成员变量也被称为"字段"
	*/

	//1.方法一.
	var p1 Person
	fmt.Println(p1)
	p1.name = "王二狗"
	p1.age = 12
	p1.sex = "男"
	p1.address = "湖北市"
	fmt.Printf("姓名:%s,年龄:%d,性别:%s,地址:%s\n", p1.name, p1.age, p1.sex, p1.address)

	//2.方法二.
	p2 := Person{}
	p2.name = "Ruby"
	p2.age = 12
	p2.sex = "女"
	p2.address = "南京市"
	fmt.Printf("姓名:%s,年龄:%d,性别；%s,地址:%s\n", p2.name, p2.age, p2.sex, p2.address)

	//3.方法三.
	p3 := Person{name: "如花", age: 23, sex: "女", address: "天津市"}
	fmt.Println(p3)

	p4 := Person{
		name:    "隔壁老王",
		age:     12,
		sex:     "男",
		address: "广西市",
	}
	fmt.Println(p4)
	//4.方法四.
	p5 := Person{name: "李小白", age: 22, sex: "女", address: "北京市"}
	fmt.Println(p5)
}

// 1.定义结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}
