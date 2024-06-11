package main

import "fmt"

func main() {
	//"" =>可解释字符串
	// ``=> 原生字符串
	var msg = "你的\\n名字是kk"
	var msg1 = `你的\n名字是kk`

	fmt.Printf("%T %s\n",msg,msg)
	fmt.Printf("%T %s\n",msg1,msg1)

	//操作
	//算数运算符：字符串连接 +
	fmt.Println(msg + msg1)

	//关系运算 > >= <= < != ==
	fmt.Println("abc"> "acd") //false
	fmt.Println("abc">= "acd") //false
	//赋值 = +=
	msg += "+赋值kk"
	fmt.Println(msg)

	//字符串定义内容必须只能为ascii码

	//索引 0 - n-1(n 字符串长度)
	msg = "abcd"
	fmt.Printf("%T,%v\n",msg[0],msg[0])
	//切片[start:end] start end -1
	fmt.Printf("%T %s\n",msg[0:2],msg[0:3])
}
