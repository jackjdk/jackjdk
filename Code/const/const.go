package main

import "fmt"

var Packagename = "package"

/*整合定义常量*/
const (
	//Packagename string = "package"
	packageMsg = "packagename"
)

func main() {

	fmt.Println(Packagename)

	const name string = "KK" /*定义的常量可以不使用*/
	/*name = "string" 常量一旦定义，不能在内部修改值*/
	const msg = "msg"
	//fmt.Println(name, packagename, packageMsg, msg)

	/*
		var llw = "KK"   #常量定义了必须使用
		fmt.Println(llw)
	*/

}
