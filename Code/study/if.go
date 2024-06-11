package main

import "fmt"

func main() {
	/*	if 布尔表达式(bool类型) {
		在布尔表达式为true时执行{}里的内容。反之为false则直接结束
		}
	*/

	//开辟一个内存，存储变量，赋值为6。if块是一个整体，取num的值(6)来和10比较，条件不成立则不执行直接结束。直接打印主函数.条件成立打印
	//执行内容并打印，最后执行主函数
	//定一个数字，如果大于10，就输出打印这个数字大于10
	/*
	num :=16

	if num >10 {
		fmt.Println("大于10")
	}

	//定一个成绩，如果大于等于60分，就打印及格
	nums :=80
	if nums >60 {
		fmt.Println("恭喜你，成绩及格")
	}



	if...else语句
		if 条件(bool类型){
		  条件成立，执行此处的A段代码。。

		} else {
		  条件不成立，执行此处的b段代码。。
	}
	注意点：
	1.if后的{,一定是和if条件写在同一行的
	2.else一定是if语句}之后,不能自己另起一行
	3.if和else中的内容,必须是二选一执行

	//给定一个成绩，如果大于等于60，则刚好及格。否则就是不及格

	score :=0   //先定义一个变量
	fmt.Println("请输入您的成绩")
	fmt.Scan(&score)  //通过键盘输入来读取score这个数值赋值给他变量
	if score >=60 {	//下面是一个完整的if~else语句块。输入大于60的数值，条件成立打印内容并且不继续执行。条件不成立则执行else语句块
		fmt.Println(score,"恭喜，你的孩子成绩及格了")
	}else{
		fmt.Println(score,"抱歉，你的孩子成绩没有达到及格线")
	}


	if 语句的嵌套:
	if 条件1{
		A段代码
	}else if 条件2{
	    B段代码
	}else if 条件3{
		C段代码
	}else{
		D段代码
	}





	//给定性别，如果是男，就去男厕所，否则去女厕所

 	sex :="泰国" 	//bool,int,string
 	if sex == "男" {
		fmt.Println("你可以去男厕所了") //如果条件成立则执行打印，并结束判断。sex 是男
	} else	if sex == "女"{						//否则向下继续执行if语句块。女，其他
		fmt.Println("你去女厕所吧")	//如果条件成立则执行并打印，结束判断。sex 是女
	} else {
		fmt.Println("我也不知道了")	//以上条件都不成立，最后会执行else并打印内容。结束
		}


	if语句的其他写法：
	if初始化语句;条件{

	}
	*/
	if num4 := 4; num4 > 0{  //这种写法只能在这个if块里使用。if结束后num4在内存中就销毁了
		fmt.Println(num4,"正数")
	}




	fmt.Println("main..over")

















/*

	// yes == "Y" "y"
	var yes string
	fmt.Print("有卖西瓜的嘛？(Y/N):")
	fmt.Scan(&yes)
	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")





	if yes == "Y" || yes == "n"{
		fmt.Println("一个西瓜")
	}
	fmt.Println("老公的想法: ")
	if yes == "Y" ||yes == "y" {
		fmt.Println("一个包子")
	} else {
		fmt.Println("十个包子")
	}

	var score int
	fmt.Print("请输入成绩:")
	fmt.Scan(&score)

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}



*/

	}



