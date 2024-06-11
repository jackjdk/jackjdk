package main

import "fmt"

func main() {
	a := 10
	fmt.Println("fun_Par()函数调用前，a:", a)
	fun_Par(a)
	fmt.Println("fun_Par()函数调用后,a:", a)

	fun_par(&a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun_three函数调用前:", arr1)
	fun_three(arr1)
	fmt.Println("fun_three()函数调用后:", arr1)

	fun_four(&arr1)
	fmt.Println("fun_four()函数调用后:", arr1)

}

func fun_four(arr2 *[4]int) {
	fmt.Println("fun_four()函数中的数组指针:", arr2)
	arr2[0] = 200
	fmt.Println("fun_four()函数中的数组指针:", arr2)
}

func fun_three(arr1 [4]int) {
	fmt.Println("fun_three()函数中数组的:", arr1)
	arr1[0] = 100
	fmt.Println("fun_three函数中修改数组:", arr1)
}

func fun_Par(num int) { //值传递:num = a =10
	fmt.Println("fun_Par()函数中,num的值:", num)
	num = 100
	fmt.Println("fun_Par()函数中修改num:", num)
}

func fun_par(p1 *int) { //传递的是a的地址，就是引用a
	fmt.Println("fun_par()函数中,p1:", *p1)
	*p1 = 200
	fmt.Println("fun_par()函数中,修改p1", *p1)
}
