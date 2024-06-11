package main

import "fmt"

func main() {
	var name string = "KK" //定义了类型，并且初始化了值
	var zerostring string  //定义变量类型，但不初始化值。
	// 初始化使用类型对应的零值(空字符串)
	var typestring = "KK" //定义变量省略类型，但不能省略初始化的值。
	// 通过对应的值类型推到变量的类型
	shortstring := "KK" //短声明(必须在函数内包含数内子块使用，不能再包级别使用)通过对应的值类型推到变量的类型
	//如果使用 var 关键字声明变量时没有使用 = 符号赋予一个初始值，那么该变量的值将会是该类型的零值

	fmt.Println(name, zerostring, typestring, shortstring)

}
