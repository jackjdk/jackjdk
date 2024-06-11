package main

import "fmt"

func main() {

	var name string = "llw"

	fmt.Println(name)

	name = "silence" //更新变量的值(赋值)
	fmt.Println(name)
	{
		//没有定义变量，是赋值.如果子块定义了var变量，打印子块的时候会优先使用定义的变量而不是赋值
		name = "aaaaa"
		var name= "ppppp"
		fmt.Println(name,name)
	}
}
