package main

import "fmt"

func main() {
	/*
		函数的返回值：
		一个函数的执行结果，返回给函数的调用处，执行结果就叫做函数的返回值
		return语句：词义“返回”
			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行
		注意点：
			A：一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处。
			B：函数返回的结果，必须和函数定义的一致：类型，个数，顺序。
			C：空白标识符， 可以使用_，来舍弃多余的返回值
			D：如果一个函数定义了有返回值，那么函数中有分支，循环，那么要保证，无论执行了哪个分支，都要有return被执行到
			E:如果一个函数没有定义返回值，那么函数中也可以使用return,专门用于结束函数的执行

	*/

	A, c, b := func5()
	fmt.Println(A, c, b)

	_, _, d := func5()
	fmt.Println(d)

}

func func5() (int, float64, string) {
	return 0, 0, "hello"
}

func ADD(e, b int) int {
	return e + b
}

func AD() {
	fmt.Println(ADD(1, 5))

}

func calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}

func calc2(a, b int) (sum, diff, product, merchant int) {

	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b

	return
}

func ma() {
	q, w, e, r := calc(9, 3)
	fmt.Println(q, w, e, r)

	fmt.Println(calc2(10, 3))
}
