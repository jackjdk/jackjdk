package main

import "fmt"

func main() {
	//len=长度,即切片初始化时的长度，没设定值默认是0
	//cap=容量(预留内存空间)
	//make 定义切片
	//append修改预留的内存空间值
	//通过append去给底层内存设定值超过指定的切片的len时，会自动创建一块新的内存空间，两个切片共用一个底层内存

	var nums []int
	fmt.Printf("%T\n", nums)
	fmt.Println(nums == nil)

	//字面量
	nums = []int{1, 2, 3}
	fmt.Printf("%#v\n", nums)
	nums = []int{1, 2, 3, 4}
	fmt.Printf("%#v\n", nums)

	//数组切片赋值
	var arrys [10]int = [10]int{1, 2, 3, 4, 5, 6} //定义的数组
	nums = arrys[1:5]                             //省略不写
	fmt.Printf("%#v\n", nums)
	//获取切片长度、容量 len cap
	fmt.Printf("%#v\n %d %d\n%", nums, len(nums), cap(nums))
	//make 函数
	nums = make([]int, 5, 5)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))
	//元素操作(增删改查)
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	nums[2] = 10
	fmt.Println(nums)
	//通过append给预留内存设定值，当设定超过内存期限值，会自动创建一块新的内存空间，新的空间会cp原来的值及内存空间是原来的2倍数，两个切片共用一个底层内存
	nums = append(nums, 1)
	fmt.Printf("%#v %d %d\n", nums, len(nums), cap(nums))

	for i := 0; i < len(nums); i++ {
		fmt.Println(i, nums[i])
	}

	//切片操作,摒弃start打印剩余的直至END，长、容量也是根据指定字符打印
	//n_cap -start
	n := nums[1:3:10]
	fmt.Printf("%T %#v %d %d\n", n, n, len(n), cap(n))

	//指定的字符限制就是长、宽这里指的是总内存(包含预留内存)，不指定宽并设置默认是nums的最大宽
	n = nums[1:3]
	//src_cap -start
	fmt.Printf("%T %v %d %d\n", n, n, len(n), cap(n))

	nums = make([]int, 4, 5) //定义一个切片，长为4，容量为5
	fmt.Println(nums)
	nums02 := nums[0:5] //:冒号表示切片的范围，0表示切片的起始索引，5表示切片的结束索引，不包含该索引处的元素，也就是到索引4结束
	// 从nums拷贝长度容量建立起nums02，包括num的容量。如果nums02没设定值的索引位置，会打印0
	nums02[1] = 1 //这里也解释下，从start开始打印，不指定0开始打印就不会全部打印，因为0则是开始。修改了1之后新建了一个内存空间共享数据
	nums02[2] = 1
	nums = append(nums, 2)
	fmt.Println(nums)

	//copy
	nums04 := []int{1, 2, 3}
	nums05 := []int{10, 20, 30, 40}
	copy(nums05, nums04) //copy需要把需要复制的数据放到后面，复制目标数据到本体中需要看双方内存对比，
	fmt.Println(nums05)

	nums05 = []int{10, 20, 30, 40} //nums05内存还有剩余，但是目标内存没有多余数据，只会复制3个
	copy(nums04, nums05)
	fmt.Println(nums04)

	//删除元素
	nums06 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums06[2:], nums06[:1]) //从前面删除.1前面的元素都舍弃
	fmt.Println(nums06[:len(nums06)-2]) //从后面删除。len返回nums06的长度之后-2。从起始索引（包括）到倒数第二个元素

	//堆栈:每次添加在队尾，移除元素在队尾 （先进后出）
	//队列：每次添加在队尾，移除元素在队头 (先进先出)。
	/*
	  queue := []int{}
	  queue = append(queue,1)
	  queue = append(queue,2)
	  queue = append(queue,3)
	  queue = append(queue,5)
	  //1,2,3,5
	  fmt.Println(queue[0])
	  queue = queue[1:]
	  //2,3,5
	  fmt.Println(queue)//每移除一个新的队头元素，(后来者居上)，然后可以继续移除队头元素

	  fmt.Println(queue[0])
	  queue = queue[1:]
	  //3,5
	  fmt.Println(queue)

	  //堆栈 从队尾删除
	  stack :=[]int{}
	  stack = append(stack,101)
	  stack = append(stack,102)
	  stack = append(stack,103)
	  stack = append(stack,104)


	  fmt.Println(stack[:len(stack)-1])
	  stack = stack[:len(stack)-1]
	  fmt.Println(stack[:len(stack)-1])
	  stack = stack[:len(stack)-1]
	  fmt.Println(stack[:len(stack)-1])
	  stack = stack[:len(stack)-1]



	  points := [][] int {}

	  point02 := make ([] int,5,10)  //make指定切片后，可以在类型后面指定切片的长宽、接着可以赋值给切片中的内存
	  point02 = []int{1,2,3,4}
	  fmt.Printf("%T %v %d %d\n",point02,point02,len(point02),cap(point02))
	 //fmt.Printf("%T\n",point02)
	 fmt.Println(point02[0:1:2])

	  points =append(points,[] int {1,2,3})
	  points =append(points,[] int {3,4,0})
	  points =append(points,[] int {3,4,0,2,4,5,6,6,6,6,6,6,6})
	  fmt.Println(points) //这里等于三个数组，不指定默认打印所有points的数组元素
	  fmt.Println(points[0])
	  fmt.Println(points[2][2])//指定索引

	  //数组是值类型
	  slice01 := [] int {1,2,3}
	  slice02 := slice01
	  slice02[0] = 10
	   fmt.Println(slice01,slice02)

	  array01 := [3]int{1,2,3}
	  array02 := array01
	  array02[0] = 10
	  fmt.Println(array01,array02)



	*/
}
