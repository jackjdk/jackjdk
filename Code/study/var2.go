package main

import "fmt"

func main() {
	/*
	var name string =  "KK"
	var msg = "hello_world"
	var desc string
	*/

	var(
		name string = "kk"
		msg 		= "hello-world"
		desc string   //声明了一个变量，但是没有值
	)
	/*
	x := "x"
	y := "y"
	*/

	x,y :="x","y"
fmt.Println(name,msg,desc,x,y)



}