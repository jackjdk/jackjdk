package main

import "fmt"

func main() {
	/*
	   	基本数据类型
	   		布尔类型：bool
	   			取值：true,false(默认为false)
	   		数值类型：
	   			整数：int
	   				有符号：最高位表示符号位，0正数，1负数，其余为表示数值
	   				int8: -128 ~ 127
	   	            int16: -32768 ~ 32767
	      				int32: -2147483648 ~ 2147483647
	      				int64: -9223372036854775808 ~ 9223372036854775807
	      				无序号：所有的位表示数值
	   				uint8: 0 ~ 255
	   				uint16: 0 ~ 65535
	   				uint32: 0 ~ 4294967295
	      	  			uint64: 0 ~ 18446744073709551615
	   			int,unit
	   			byte:unit8
	   			rune:int32
	   		浮点：生活中的小数
	   			float32，float64
	   		复数：complex，
	   	字符串：string
	   2.复合数据类型
	   	array、slice、map、function、pointer、struct、interface、channel

	   	数据类型转换:type convert
	   	go语言是静态语言，定义，赋值，运算必须类型一致

	   	语法格式:Type(value)
	   	注意点：兼容类型可以转换
	   	常数：在有需要的时候，自动转型
	   	变量：需要手动转型
	*/
	var a int8
	a = 10

	var b int16
	b = int16(a) //a的值赋给了b
	fmt.Println(a, b)

	f1 := 4.13
	var c int
	c = int(f1) //类型转换的语法是在要转换的类型名称前面加上括号并指定要转换的值,例如c = int(f1)。这将把浮点数f1转换为整数类型，并将结果赋值给c
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)
	sum := f1 + 100 //常数再有需要的时候自动转型
	fmt.Println("%T,%f\n", sum, sum)

	//float类型---------------------------------------------------------
	//var hight float32 = 1.68 //指定类型
	//
	//var hightType = 1.68 //类型推断
	//
	//fmt.Printf("%T %#v %f\n", hight, hight, hight)
	//fmt.Printf("%T %v %f\n", hightType, hightType, hightType)

	//var k = 1e3
	//fmt.Println(k)
	////操作
	////算术运算： + - * / ++ --
	//var (
	//	f1 = 1.2
	//	f2 = 2.36
	//)
	//fmt.Println(f1 + f2)
	//fmt.Println(f1 - f2)
	//fmt.Println(f1 * f2)
	//fmt.Println(f1 / f2)
	//
	//f1++ //2.2
	//f2-- //1.35
	//fmt.Println(f1, f2)
	//
	////关系运算 > < >= <=  如果想要判断!= ==  判断差值在一定区间范围内
	//fmt.Println(f1 > f2)
	//fmt.Println(f1 >= f2)
	//fmt.Println(f1 < f2)
	//fmt.Println(f1 <= f2)
	//
	////赋值运算 = += -= /= *=
	//f1 = 1.32
	//f1 += f2 //f1 = f1 + f2
	//fmt.Println(f1)

	//string_Type ----------------------------------------------------------------------
	//"" =>可解释字符串
	//``=> 原生字符串
	var msg = "你的\\n名字是kk"
	var msg1 = `你的\n名字是kk`

	fmt.Printf("%T %s\n", msg, msg)
	fmt.Printf("%T %s\n", msg1, msg1)

	//操作
	//算数运算符：字符串连接 +
	fmt.Println(msg + msg1)

	//关系运算 > >= <= < != ==
	fmt.Println("abc" > "acd")  //false  根据Unicode字典排序
	fmt.Println("abc" >= "acd") //false
	//赋值 = +=
	msg += "+赋值kk"
	fmt.Println(msg)

	//字符串定义内容必须只能为ascii码
	//
	//索引 0 - n-1(n 字符串长度)
	msg = "abcd"
	fmt.Printf("%T,%v\n", msg[0], msg[0]) //输出格式根据切片的内容输出值
	//切片[start:end] start end -1
	fmt.Printf("%T %s\n", msg[0:2], msg[0:3])

}
