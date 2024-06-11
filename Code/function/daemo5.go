package main

import "fmt"

func main() {
	/*
	递归函数(recursion):一个函数自己调用自己，就叫做递归函数
			递归函数要有一个出口，逐渐的向出口靠近
	 */

	//1.求1-5的和
	sum := getSum1(5)
	fmt.Println(sum)

}

func getSum1(n int)int {
	if n == 1 {
		return 1
	}
	fmt.Println("计算f(n):",n)
	return getSum1(n-1) + n

}

/*
getSum（5）
	getSum(4) + 5
		getSum(3) + 4
			getSum(2) + 3
				getSum(1) + 2
				1

 */