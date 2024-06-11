package main

import (
	"fmt"
)

func main() {
	const (
		A = "test"
		B //使用前一个常量的值进行初始化(B)
		C //使用前一个常量的值进行初始化(C=>B)
		D = "testD"
		E //使用前一个常量的值进行初始化(D)
		F //使用前一个常量的值进行初始化F=>E)
	)
	fmt.Println(A, B, C, D, E, F)

}
