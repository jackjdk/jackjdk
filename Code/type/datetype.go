package main

import (
	"fmt"
)

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
	a =  10

	var b int16
	b = int16(a)
	fmt.Println(a,b)

	f1 :=4.13
	var c int
	c = int(f1)
	fmt.Println(f1,c)

	f1 = float64(a)
	fmt.Println(f1,a)
}
