package main

import "fmt"

func main() {

	//算术运算符：+，-，*，/,%,++,--
	//+
	//-
	//*:乘法
	///：取商，两个数相除，取商
	//%：取余，取模，两个数相除，取余数
	//++：给自己加1 (i++)
	//--:给自己减1  (i--)

	//a := 10
	//b := 3
	//div := a / b //取商
	//mod := a % b //取余
	//fmt.Printf("%d / %d = %d\n", a, b, div)
	//fmt.Printf("%d %% %d = %d\n", a, b, mod)
	//
	//c := 3
	//c++ //就是给c加1  使用for循环指定次数的话每次回让c++加1
	//fmt.Println(c)
	//
	//c-- //给c减1
	//fmt.Println(c)

	//关系运算符：> ,<,>=,<=,==,!=   -------------------------------------------------------
	// 结果总是bool类型的：true，false
	//==：表示比较两个数值是相等的
	//!=：表示比较两个数值是不想等的

	//var a, b, c = 3, 5, 3
	//
	//res := a == b
	//fmt.Println(res)
	//
	//res2 := a == c
	//fmt.Println(res2)
	//
	//fmt.Println(a != b, a != c)

	//逻辑运算符：操作数必须是bool，运算结果也是bool------------------------------------------------------------
	//		逻辑与：&&
	//			运算规则：所有的操作数都是真，结果才为真，有一个为假，结果就为假
	//					"一假则假，全真才真"
	//		逻辑或：||
	//			运算规则：所有的操作数都是假，结果才为假，有一个为真，结果就为真
	//					"一真则真，全假才假"
	//		逻辑非： !
	//			!True-->false //取反
	//			!False-->true

	//Q1 := true
	//Q2 := false
	//Q3 := true
	//
	//res1 := Q1 && Q2
	//fmt.Printf("res1: %t\n", res1)
	//fmt.Printf("res1: %t\n", res1)
	//
	//res2 := Q1 && Q2 && Q3
	//
	//fmt.Printf("res2: %t\n", res2)
	//
	//res3 := Q1 || Q3
	//res4 := Q1 || Q2
	//fmt.Printf("res3: %t\n", res3)
	//
	//fmt.Printf("res3: %t\n", res4)
	//
	//fmt.Printf("Q1:%t,!Q1:%t\n", Q1, !Q1) //根据值取反
	//fmt.Printf("Q2:%t,!Q2:%t\n", Q2, !Q2)

	/*
			var age int8 = 126
			var age1 int = 31
			fmt.Printf("%T,%#v,%d\n",age,age,age)
			fmt.Printf("%T,%#v,%d\n",age1,age1,age1)
			了解
			8进制：0???<8
			16进制 0x?? ? < 16 0-9A-F
			070 => 10进制：56
			078 => 10进制: 76
		    二进制存储
			base2
			7=> 0111=> 1*2^2+1*2^1+1*2^0

			fmt.Println(070,78)
			//常用操作
			//算术运算 + - * / % ++ --
			//a = b*c +d => a % c = d
			a,b := 2 , 4
			fmt.Println(a + b)
			fmt.Println(a % b) //2  取余数
			fmt.Println(a / b) //0	取商

			a++
			b--
			fmt.Println(a,b)//3,3

			//常用操作
			//关系运算 > < >= <= != ==
			fmt.Println(a > b ) //false
			fmt.Println(a < b ) //false
			fmt.Println(a >= b ) //true
			fmt.Println(a <= b ) //true
			fmt.Println(a != b ) //false
			fmt.Println(a == b ) //true

			//仅了解,位运算
			//7 => 0111
			//2 => 0010
			按位与: & 两个都为1结果为1， 0010 7&2 =>2
			按位或: |,只要有一个为1结果为1, 0111 7|2 => 7
			取非: ^,对于有符号位来说首字符不变，1=>0,0=>1


		var (
			i	int = 1
			i32 int32 = 1
			i64 int64 = 1
		)
		//类型转换 type(value) int32(i) int(i32) int64(i) int(i64)
			fmt.Printf("%T\n",i + int(i32))
			fmt.Printf("%T\n",i + int(i64))
			fmt.Printf("%T\n",int32(i) + i32)
			fmt.Printf("%T\n",int64(i) + i64)
			//%T	输出值的类型
			//%v	使用默认格式输出值，或者如果方法存在，则使用类性值的String()方法输出自定义值
			//%s	输出以原生的UTF8字节表示的字符，如果console不支持utf8编码，则会乱码
			//%U	一个Unicode表示法表示的整型码值
			//%X	大写的十六进制数值
			//%x	小写的十六进制数值
			//%o	八进制整数
			//%d	十进制整数
			//%c	一个Unicode的字符
			//%b	一个二进制整数，将一个整数格式转化为二进制的表达方式

			var(
				achar		byte = 'A'
				aint		byte = 65
				unicodePoint rune = '中'
				//字符串 => 内存 （01） => 转换 => 编码（utf8m,utf16编码，gbk，gb2312）

			)
		fmt.Println(achar,aint)
		fmt.Println(unicodePoint)

		fmt.Printf("%d %b %o %x %U %c %c",achar,15,15,15,unicodePoint,achar,65)
	*/

	//赋值运算符
	//=,+=,-=,*=,/=,%=,<<=,>>=,&=,|=,^=
	//= 把右侧的数值，赋值给=左边的变量
	//+=， a += b，相当于 a = a + b
	//a++,a += 1

	var a int
	a = 3
	fmt.Println(a) //3

	a += 4         //a = a + 4
	fmt.Println(a) //7

	a -= 3
	fmt.Println(a) //4

	a /= 3
	fmt.Println(a) //2

	a %= 1
	fmt.Println(a) //0

}
