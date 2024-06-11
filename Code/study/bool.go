package main

import "fmt"

func main() {
	//布尔类型
	isGirl :=false

	fmt.Printf("%T,%#v",isGirl,isGirl)
	//%T	输出值的类型
	//%v	使用默认格式输出值，或者如果方法存在，则使用类性值的String()方法输出自定义值

	//操作
	//逻辑运算
	a,b,c,d := true,true,false,false

	//与:左操作书与右操作数都为true，结果为true
	fmt.Println("a,b:",a && b) // true && true :true
	fmt.Println("a,c:",a && c) // true && false :false
	fmt.Println("c,b:",c && b) // false && true :false
	fmt.Println("c,d:",c && d) // false && false :false
	//或：左操作数与右操作数只要有一个为true，结果为true
	fmt.Println("a,b:",a || b) // true && true :true
	fmt.Println("a,c:",a || c) // true && true :true
	fmt.Println("c,b:",c || b) // true && true :true
	fmt.Println("c,d:",c || d) // true && true :false
	//非：取反 true=> false, fasel=> true
	fmt.Println("a:", !a) // !true :false
	fmt.Println("c:", !c) // !false :true
	//关系
	fmt.Println(a == b) //true == true :true
	fmt.Println(a != c) //true == false :true
	fmt.Println(a == c) //true == false : false
	fmt.Println(c != b) // false == true : true


	fmt.Println("%t,%t",a,c)

	var zero bool //零值默认是false
	isBoy := true
	isGirll := false
	fmt.Println(zero,isBoy,isGirll)
}
