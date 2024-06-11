package main

import "fmt"

func main() {
	const (
		Mon= iota //在常量组内使用iota初始化0，每次调用+1
		Tuesd
		Wed
		Thur
		Fir
		Sat
		Sun
	)
fmt.Println(Mon,Tuesd,Wed,Thur,Fir,Sat,Sun)
}
