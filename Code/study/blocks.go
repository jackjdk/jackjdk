package main

import "fmt"

func main() {
	//作用域：定义标识符可以使用的范围
	//在Go中用{}来定义作用域的范围
	//使用原则：子语句块可以使用父语句块中的标识符，父不能使用子的
	//使用原则：子块在使用父的标识符并且修改重新在子块中修改标识符并使用
	//子块级别定义的内部变量仅供子块使用，还可以使用外部定义的变量。外部使用不了子块(内部)定义的变量。go执行模块是由下而上
	outer := 1
	{
		outer := 22
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		{
			inner1 := 3
			outer :=4
			fmt.Println(outer,inner,inner1,outer)
		}
	}
	//fmt.Println(inner,outer)
	fmt.Println(outer)
}
