package main

import "fmt"

// fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入
func main() {
	//var name = "KK"
	//var age int= 12
	//fmt.Println(age,name)//打印一行值+空格输出+换行
	//fmt.Println(name)
	//fmt.Print(name)   //原样输出，打印一行变量不加换行
	//fmt.Printf("%T,%V,%#V",name,name,name) //格式输出 %T打印字符串的类型  %V打印字符串的变量

	/*
		%v：根据变量类型，输出该变量的值。
		%T：输出变量的类型。
		%t：输出布尔类型变量的值，true或false。
		%d：输出十进制整数。
		%b：输出二进制整数。
		%o：输出八进制整数。
		%x：输出十六进制整数，使用小写字母。
		%X：输出十六进制整数，使用大写字母。
		%f：输出浮点数，保留小数点后6位。
		%e：输出科学计数法表示的浮点数。
		%s：输出字符串。
		%p：输出指针地址。
		%c：输出字符。"
	*/

	var name string
	fmt.Print(&name)
	fmt.Scan(&name)
	fmt.Println("名字:", name)

	var age int
	fmt.Print("请输入你的年龄")
	fmt.Scan(&age)
	fmt.Println("年龄:", age)

	var height float64
	fmt.Print("请输入你的身高")
	fmt.Scan(&height)
	fmt.Println("身高:", height)

	var weight float64
	fmt.Print("请输入你的体重")
	fmt.Scan(&weight)
	fmt.Println("体重:", weight)
}
