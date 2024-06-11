package main

import "fmt"

func main() {
/*
	函数：function
	一.概念：具有特定功能的代码，可以被多次调用执行
	二.意义
		1.可以避免重复的代码
		2.增强程序的扩展性
	三.使用
		step1：函数的定义，也叫声明
		step2：函数的调用，就是执行函数中代码的过程
	四.语法
		1.定义函数的语法
			func funcName(parametername type1，parametername type2) (output1 type1，output2 type2){
				//这里是处理逻辑代码
				//返回多个值
				return value1 value2
		}
			A：func，定义函数的关键字
			B：funcname，函数的名字
			C：()，函数的标志
			D:参数列表：形式参数用于接受外部传入函数中的数据
			E：返回值列表：函数执行后返回给调用处的结果
		2.调用函数的语法
			函数名(实际参数)
		3.注意事项
			A:函数必须先定义，在调用，如果不定义：undefined：getsum
 			定义了函数，没有调用，那么函数就失去了意义

			B：函数名不能冲突
			C：main（）是一个特殊的函数，作为程序的入口，由系统自动调用
					而其他函数，程序中通过函数名来调用

		函数的参数：
			形式参数：也叫形参。函数定义的时候，用于接受外部传入的数据的变量。
					函数中，某些变量的数值无法确定，需要由外部传入数据
			实际参数：也叫实参。函数调用的时候，给形参复制的实际的数据
		函数调用：
			1：函数名：声明的函数名和调用的函数名要统一
			2.实参必须严格匹配形参:顺序，个数，类型，一一对应的
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
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处，执行结果就叫做函数的返回值
		return语句：
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处。
			函数返回的结果，必须和函数定义的一致：类型，个数，顺序。

			1.将函数的结果返回给调用处
			2.同时结束了该函数的执行

	空白标识符，专门用于舍弃数据：_


 */
	//求1-10的和
	getsum(20)	//函数的调用处

	//求1-100的和
	getsum(100)


	//求1-20的和
	getsum(10)

	//4.求2个整数的和
	getAdd(20,20)
	fun1(1.3,2.4,"a")

	//5.求总和
	getSum(1,2,3,4,5 )

	//6.切片
	s1 :=[]int{1,2,3,4,5}
	getSum(s1...)

	//返回值
	rt := AB(1, 5)   //定义的函数返回值要赋值给一个变量，最后输出变量的值
	fmt.Println(rt)



	res1,res2 := rectangle(5,3)
	fmt.Println("周长：",res1,"面积:",res2)
	res3,res4 := rectangle2(5,3)
	fmt.Println("周长:",res3,"面积:",res4)

	_,res6 :=rectangle (100,3)
	fmt.Println(res6)

}


//定义一个函数，用于求1-10的和
func getsum(n int)  {
	sum := 0

	for i := 1; i <= n; i++{
		sum += i
	}
	fmt.Printf("1-%d的和是: %d\n",n,sum)
}


func getAdd(a,b int)int{
	sum := a + b
	fmt.Printf("%d + %d = %d \n",a,b,sum)
	return a +b
}


func  fun1(a,b float64,c string){
	fmt.Printf("a:%2f,b:%2f,c:%s\n",a,b,c)
}

func getSum(nums...int){
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println("总和是:",sum)
}

func AB(e,b int)int {
	return e + b
}




//函数，用于求矩形的周长和面积
func  rectangle(len,wid float64) (float64,float64) {
	perimeter := (len +wid) *2
	area := len * wid
	return perimeter,area
}

func rectangle2(len,wid float64)(peri float64,area float64){
	peri = (len +wid)*2
	area = len*wid
	return
}

func fun4()(float64,float64,string){
	return 2.4,5.9,"hello"
}

