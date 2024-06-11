package main

import "fmt"

func main() {
	var age = 30

	age = age +1
	fmt.Println("明年:",age)
	age = age +1
	fmt.Println("后年:",age)

	fmt.Println("打印第一行")  //打印一行变量加换行
	fmt.Println("打印第二行")
	fmt.Print("打印第一行")	//打印一行变量不加换行
	fmt.Print("打印第二行")
	fmt.Printf("%T,%S","%d",1,"KK","kk") // %T打印字符串的类型  %V打印字符串的变量
}
