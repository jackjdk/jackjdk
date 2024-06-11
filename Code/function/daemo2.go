package main

import "fmt"

func main() {
	/*
	可变的参数：
				概念：一个函数的参数的类型确定，但是个数不确定，就可以使用可变参数。

				语法：
					参数名：...参数的类型
					对于函数，可变参数相当于一个切片
					调用函数的时候，可以传入0-多个参数
					println(),printf(),print()
					append
				注意事项：
					A:如果一个函数的参数是可变参数，同时还有其他的参数，可变参数要放在参数列表的最后
					B:一个函数的参数列表中最多只能有一个可变参数

	 */
	//1.求和
	getSUM()
	getSUM(1,2,3,4,5)
	getSUM(1,2,3,4,5,6,7,8,9,10)

	//2.切片
	s1 := []int{1,2,3,4,5}
	getSUM(s1...)




	nums06 := [] int {1,3,5,8}
	fmt.Println(nums06[:1])
	fmt.Println(nums06[2:])
	nums06 = append(nums06[:1],nums06[2:]...)
	fmt.Println(nums06)
}

func getSUM(nums ... int){
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]

	}
	fmt.Println("总和是:",sum)
}

