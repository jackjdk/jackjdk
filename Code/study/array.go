package main

import (
	"fmt"
)

// 数组
// 数组必须定义长度

func main() {
	var num [10]int
	var num1 [5]bool
	var num2 [5]string
	var nums [5]int
	fmt.Printf("%T", num)
	fmt.Println(num)
	fmt.Println(num1)
	fmt.Println(num2)
	//字面量
	num = [10]int{10, 20, 30}
	fmt.Println(num)
	// 根据索引设定数组里指定位置的数字
	num = [10]int{0: 1, 9: 9}
	fmt.Println(num)
	//推到，意味着元素有多少个数组，数组需要全部用到，否则有问题
	nums = [...]int{1, 2, 3, 1, 2}
	fmt.Println(nums)
	//使用短声明，没有赋值的数组，会打印0
	nums2 := [5]int{1, 20}
	fmt.Printf("%T %#v", nums2, nums2)
	//使用短声明，并且结合索引推到，数组可以任意使用数量
	nums3 := [...]int{1, 2}
	fmt.Printf("%T %#v\n", nums3, nums3)
	//使用短声明变量给元素赋值，也可以指定给某个位置的元素赋值，没有赋值的皆为0
	nums4 := [15]int{0: 10, 3: 20, 14: 666}
	fmt.Printf("%T %#v\n", nums4, nums4)

	//操作
	//运算数量需要对等，否则有问题
	nums5 := [3]int{1, 2, 3}
	nums6 := [3]int{1, 2, 3}
	fmt.Println(nums5 == nums6)
	fmt.Println(nums5 != nums6)
	//通过len获取数组的长度
	fmt.Println(len(nums4))
	// 取数组里面指定的值，指定变量+数组位置 。索引 0,1,2 ...len(array)
	fmt.Println(nums4[0], nums4[3])
	//指定变量的数组以及数组位置，并修改数组的值
	nums4[0] = 1000
	fmt.Println(nums4)
	//通过 for len 去遍历数组,数组里提前设置好的值也会打印
	for i := 0; i < len(nums4); i++ {
		fmt.Println(i, ":", nums4[i])
	}
	//简短声明
	for index, value := range nums4 {
		fmt.Println(index, ":", value)

	}
	//简短声明
	for _, value := range nums4 {
		fmt.Println(value)
	}

	var value int
	_, value = 3, 6
	fmt.Println(value)
	/*
	   //通过简短声明给变量赋值
	   	var value int
	   	_,value,o := 3,6,9
	   	fmt.Println(value,o)
	*/

	//切片,start开始的位置和END结束的位置，取END位置的字符
	//s := "123123123"
	fmt.Printf("%T %#v", nums4, nums4[1:15:15])
	//多维数组在创建时自动初始化为0值，需要赋值
	var marrays [4][2]int
	marrays = [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Printf("%T %v\n", marrays, marrays)
	marrays[0] = [2]int{100, 101}
	fmt.Println(marrays)
	marrays = [4][2]int{0: {0: 101, 1: 102}}
	fmt.Println(marrays)

}
