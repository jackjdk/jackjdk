package main

import "fmt"

func main() {
	/*
		    OOP中的继承性
				如果两个类(class)存在继承关系，其中一个是子类，另一个作为父类，那么：
				1.子类可以直接访问父类的属性和方法
				2.子类可以新增自己的属性和方法
				3.子类可以重写父类的方法(orverride，就是将父类已有的方法，重新实现)

			Go语言的结构体浅谈
				1.模拟继承性
					type A struct{
						field
					}
					type B struct{
						A//匿名字段
					}

				2.模拟聚合关系:has - a
					type C struct{
						field
					}
					type D struct{
						c C //聚合关系
					}
	*/
	//1.创建Person7类型
	p1 := Person7{name: "王二狗", age: 31}
	fmt.Println(p1.name, p1.age) //父类对象，访问父类的字段属性
	p1.eat()                     //父类对象，访问父类的方法

	//2.继承Student7类型
	s3 := Person7{"Jake", 21}
	fmt.Println(s3.name, s3.age)
	s1 := Student7{Person7{"Jake", 31}, "东营科技大学"}
	fmt.Println(s1.name, s1.age) //s1.Person7.name
	fmt.Println(s1.age)          //子类对象，可以直接访问父类的字段。(其实就是提升字段)
	fmt.Println(s1.school)       //子类对象，访问自己新增的字段属性
	fmt.Println(s1.study)
	fmt.Println(s1.eat)

	s1.eat()   //子类对象,访问父类的方法
	s1.study() //子类对象,访问自己新增的方法
	s1.eat()   //如果存在方法的重写，子类对象访问重写的方法

}

// 1. 定义一个"父类"结构体
type Person7 struct {
	name string
	age  int
}

// 2.定义一个"子类"结构体
type Student7 struct {
	Person7 //结构体嵌套，模拟继承性
	school  string
}

// 3.定义两个一个Person7的方法
func (p Person7) eat() {
	fmt.Println("父类的方法,吃窝窝头")
}

// 4.新增方法
func (s Student7) study() {
	fmt.Println("子类新增的方法,学生学习啦")
}

func (s Student7) eat() {
	fmt.Println("子类重写的方法:吃炸鸡和啤酒")
}
