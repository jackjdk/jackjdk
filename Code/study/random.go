package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
/*

 */
	num1 := rand.Int()
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num :=rand.Intn(10)
		fmt.Println(num)
	}
	rand.Seed(1000)
	num2 := rand.Intn(10)
	fmt.Println("-->",num2)

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n",t1)
	//时间戳：指定时间，距离1970年1月1日0点0分0秒 之间的时间差值：秒 纳秒
	timeStamp1:=t1.Unix() //秒
	fmt.Println(timeStamp1)

	timeStamp2:=t1.UnixNano()
	fmt.Println(timeStamp2)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println("-->",rand.Intn(100))
	}

}
