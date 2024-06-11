package main

import "fmt"

// var Lvliangwei string = "test"  //全局变量可以跨包去调用
func main() {
	/* var变量的定义
	   变量: variable
	   概念: 一小块内存，用于存储数据，在程序运行过程中数值可以改变

	   使用：
	   	steep1：变量的声明，也叫定义
			第一种 var 变量名 数据类型
					变量名 = 赋值
			第二种 简短声明，省略var
					变量名 := 赋值
			第三种 类型推断，省略数据类型
				var 变量名 = 赋值
	   	steep2：变量的访问，赋值和取值


	   Go的特性
	   	静态语言：强类型语言
	   		Go，java，C++,C#。。
	   	动态语言：弱类型语言
	   		JavaScript，php，python，Ruby。。

	var name = "KK"
	var age int= 12
	fmt.Println(age,name)//打印一行值+空格输出+换行
	fmt.Println(name)
	fmt.Print(name)   //原样输出，打印一行变量不加换行
	fmt.Printf("%T,%V,%#V",name,name,name) //格式输出 %T打印字符串的类型  %V打印字符串的变量

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
	//变量定义了必须被使用
	//第一种：定义变量，然后进行赋值
	//var num1 int
	//num1 = 30
	//fmt.Println("num1的数值是:%d\n", num1)
	////第二种：写在一行
	//var num2 int = 15
	//fmt.Println("num2的数值是：%d\n", num2)
	////第三种：类型推断
	//var name = "小静"
	//fmt.Println("类型是：%T，数值是：%d\n", name)
	//
	////多个变量的同时定义
	//var a, b, c int
	//a = 1
	//b = 2
	//c = 3
	//fmt.Println(a, b, c)

	//var a = 20
	//a = 21 //重新为a赋值
	//fmt.Println(a)

	//简短声明 在同一个作用域中，已存在同名的变量，则之后的声明初始化，则退化为赋值操作。但这个前提是，最少要有一个新的变量被定义，且在同一作用域，例如，下面的y就是新定义的变量

	//x := 140
	//fmt.Println(&x)
	//x, y := 200, "abc"
	//fmt.Println(&x, x)
	//fmt.Print(y)

	//常量的使用---------------------------------------------------------------------------------

	/*整合定义常量*/

	const (
		packagename string = "package"     //显试类型定义
		packageMsg         = "packagename" //隐试类型定义
	)

	const name2 = "KK" /*定义的常量可以不使用*/
	/*name = "string" 常量一旦定义，不能在程序运行过程的时候修改值*/
	const msg = "msg"
	fmt.Println("输出类型:%T,数据是:\n", packagename, packageMsg)

	//
	//const (
	//	A = "test"
	//	B //使用前一个常量的值进行初始化(B)
	//	C //使用前一个常量的值进行初始化(C=>B)
	//	D = "testD"
	//	E //使用前一个常量的值进行初始化(D)
	//	F //使用前一个常量的值进行初始化F=>E)
	//)
	//fmt.Println(A, B, C, D, E, F)
	//
	//const Name string = "KK"
	////省略类型
	//const PI = 3.1415926
	////定义多个常量，(类型相同)
	//const C1, C2 int = 1, 2
	////定义多个常量,（类型不相同）
	//const (
	//	C3 string = "silence"
	//	C4 int    = 1
	//)
	////定义多个常量，省略类型
	//const C5, C6 = "silence", 1
	////const必须设置默认值
	////const C7 string
	//fmt.Println(Name)
	//fmt.Println(PI)
	//fmt.Println(C1, C2)
	//fmt.Println(C3, C4)
	//fmt.Println(C5, C6)
	////fmt.Println(C7)
	//const (
	//	C7 int = 1
	//	C8
	//	C9
	//	C10 float64 = 3.14
	//	C11
	//	C12 string = "L"
	//)
	//fmt.Println(C7, C8, C9, C10, C11, C12)

	//iota------------特殊常量iota------------------------------------------
	//枚举，const+iota

	// 错误示例代码

	//const (
	//	a
	//	b
	//	c = iota
	//	d
	//	e
	//)
	//fmt.Println(a, b, c, d, e)

	//编译时产生的错误信息：
	//# command-line-arguments

	//从上边的示例代码中可以得知
	//iota 并不会给括号作用域范围内使用 iota 赋值的那个常量之前的常量赋值
	//只会给括号作用域内使用 iota 初始化的那个常量后边的所有常量自增 1 后赋值。

	const (
		E1 = 10 + iota //iota 默认初始值为 0，我们可以定义的常量初始值从 10 开始
		E2
		E3
	)
	fmt.Println(E1)
	fmt.Println(E2)
	fmt.Println(E3)

	//作用域：定义标识符可以使用的范围
	//在Go中用{}来定义作用域的范围
	//使用原则：子语句块可以使用父语句块中的标识符，父不能使用子的
	//使用原则：子块在使用父的标识符并且修改重新在子块中修改标识符并使用
	//子块级别定义的内部变量仅供子块使用，还可以使用外部定义的变量。外部使用不了子块(内部)定义的变量。go执行模块是由下而上
	//outer := 1
	//{
	//	outer := 22
	//	inner := 2
	//	fmt.Println(outer)
	//	fmt.Println(inner)
	//	{
	//		inner1 := 3
	//		outer :=4
	//		fmt.Println(outer,inner,inner1,outer)
	//	}
	//}
	////fmt.Println(inner,outer)
	//fmt.Println(outer)

}
