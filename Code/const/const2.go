package main

import (
	"fmt"
)

func main() {
	const Name string = "KK"
	//省略类型
	const PI = 3.1415926
	//定义多个常量，(类型相同)
	const C1, C2 int = 1, 2
	//定义多个常量,（类型不相同）
	const (
		C3 string = "silence"
		C4 int    = 1
	)
	//定义多个常量，省略类型
	const C5, C6 = "silence", 1
	//const必须设置默认值
	//const C7 string
	fmt.Println(Name)
	fmt.Println(PI)
	fmt.Println(C1, C2)
	fmt.Println(C3, C4)
	fmt.Println(C5, C6)
	//fmt.Println(C7)
	const (
		C7 int = 1
		C8
		C9
		C10 float64 = 3.14
		C11
		C12 string = "L"
	)
	fmt.Println(C7, C8, C9, C10, C11, C12)
	//枚举，const+iota
	const (
		E1 int = iota
		E2
		E3
	)
	fmt.Println(E1, E2, E3)

}
