package main

import "fmt"

func main() {
	/*
	函数指针：一个指针，指向了一个函数的指针
		因为Go语言中，function，默认看作一个指针，没有*

		slice，man，function
	指针函数：一个函数，该函数的返回值是一个指针

	 */

	var a func()
	a = fun1
	a()

}

func fun1(){
	fmt.Println("fun1()...")
}