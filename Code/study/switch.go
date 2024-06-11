package main

import "fmt"

func main() {
	/*
	switch语句：
	语法结构：
		switch 变量名{
		case 数值1：分支1
		case 数值2：分支2
		case 数值3：分支3
		default:
			最后一个分支
	}
	注意事项：
	1.switch可以作用在其他类型上，case后的数值必须和switch作用的变量类型一致。提示：变量即使写的是int用""也可以将int当做字符串使用，但是会不匹配，所以不可用
	2.case是无序的
·	3.case后的数值是唯一的
	4.default是可选的操作

	num := 1
	switch num {
	case 3:
		fmt.Println("第三季度")
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 4:
		fmt.Println("第四季度")
	}


	//模拟计算器
	num1 := 0
	num2 := 0
	oper := ""
	fmt.Println("请输入一个整数")
	fmt.Scan(&num1)
	fmt.Println("请再次输入一个整数")
	fmt.Scan(&num2)
	fmt.Println("请输入一个操作：+,-,*,/")
	fmt.Scan(&oper)
	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d\n",num1,num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n",num1,num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n",num1,num2, num1*num2)
	case"/":
		fmt.Printf("%d / %d = %d\n",num1,num2, num1/num2)
	}

	1.switch的标注写法:
	switch 变量{
	case  数值1：分支1
	case  数值2：分支2
	。。。。。
	defalut：
	最后一个分支
	}
	2.省略switch后的变量，相当于直接作用在true上
	switch{ true
	case true:
	case false:
	}


	switch{
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	成绩:
	[0-59],不及格
	[60,69],及格
	[70,79],中
	[80,89],良好
	[90,100],优秀

	score := 88
	switch  {
	case score >= 0 && score < 60:
		fmt.Println(score,"不及格")
	case score >= 60 && score < 70:
		fmt.Println(score,"及格")
	case score >= 70 && score < 80:
		fmt.Println(score,"中等")
	case score >= 80 && score < 90:
		fmt.Println(score,"良好")
	case score >= 90 && score <= 100:
		fmt.Println(score,"优秀")
	default:
		fmt.Println("成绩有误")
	}
	fmt.Println("main....over")


fmt.Println("")
	letter := "Y"
	switch letter{
	case "A","B","E","Y":
		fmt.Println(letter,"是元音..")
	case "M","N":
		fmt.Println("M或N")
	default:
		fmt.Println("其他")
	}



	一个月的天数
	1，3,5,7,8,10,12
		31
	4,6,9,11
		30
	2:29/28

 	month := 5
 	day := 0
 	year := 2000
	switch month {
	case 1,3,5,7,8,10,12:
		day = 31

	case 4,6,9,11:
		day = 30
	case 2:
		if year%400 == 0 || year%4 == 0 && year%100 != 0 {
			day = 29
		} else {
			day = 28
		}
	default:
		fmt.Println("月份有误")
	}
	fmt.Printf("%d 年 %d 月 的天数量: %d\n",year,month,day)
	fmt.Println("---")


	*/


 	switch  language := "golang";language { //在switch里定义变量
	case "golang":
		fmt.Println("Go语言")
	case "JAVA语言":
 		fmt.Println("JAVA语言")
	case "Python语言":
		fmt.Println("Python语言")
	default:
		fmt.Println(language)
	}











/*
	var yes string
	fmt.Print("有卖西瓜的嘛？(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法:")
	fmt.Println("十个包子")
	switch yes {
	case "y","Y":
		fmt.Println("买一个西瓜")
	}

	fmt.Println("老公的想法:")

	switch yes {
	case "y","Y":
		fmt.Println("买一个包子")
	default:
		fmt.Println("买十个包子")
	}

	var score int
	fmt.Print("请输入成绩:")
	fmt.Scan(&score)

	switch  {
	case score >= 90:
	 	fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("E")
	}


	 */
}
