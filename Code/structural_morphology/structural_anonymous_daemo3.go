package main

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段：
		匿名结构体：没有名字的结构体，
			在创建匿名结构体时，同时创建对象
			变量名 := struct{
				定义字段Field
		}{
			字段进行赋值
		}

		匿名字段：一个结构体的字段没有字段名

		匿名函数：

	*/

	s1 := Student{name: "张三", age: 18}
	fmt.Println(s1.name, s1.age)

	func() {
		fmt.Println("hello world...")
	}()

	s2 := struct {
		name string
		age  int
	}{
		name: "李四",
		age:  14,
	}
	fmt.Println(s2.name, s2.age)

	//w1 := Worker1{name: "王二狗", age: 30}
	//fmt.Println("打印Worker1", w1.name, w1.age)

	w2 := Worker3{"李小花", 32}
	fmt.Println(w2)
	//fmt.Println(w2.string)
	//fmt.Println(w2.int)
}

type Worker3 struct {
	name string
	age  int
	//string //匿名字段
	//int    //匿名字段，默认使用数据类型作为名字。那么匿名字段的类型就不能重复，否则会冲突
}

type Worker1 struct {
	name string
	age  int
}

type Student struct {
	name string
	age  int
}
