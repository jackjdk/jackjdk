package main

import (
	"bytes"
	"fmt"
)


func main() {
	var bytes01 [] byte = []byte {'a','b','c','z','s'} //通过长度和输出长度来限制字节的定义
	fmt.Printf("%T","%v\n",bytes01,bytes01)


	string := string(bytes01)
	fmt.Printf("%T %v\n",string,string)

	bs := []byte(string)
	fmt.Printf("%T,%v\n",bs,bs)
	fmt.Println("\"bs\"")
	fmt.Println("I'mzq")
	//字节数组可以通过将其转换为字符串或者将字符串转换为字节数组来进行类型转换。但需要注意的是，转换的过程中要确保数据的有效性，
	//特别是在涉及到编码的情况下。
	//int类型使用{} string类型使用[""]
	fmt.Println(bytes.Compare([]byte("abc"),[]byte("def")))
	//Compare函数返回一个整数表示两个[]byte切片按字典序比较的结果。如果a==b返回0；如果a<b返回-1；否则返回+1。nil参数视为空切片。
	fmt.Println(bytes.Compare([]byte{2},[]byte{1}))
	fmt.Println(bytes.Compare([]byte{1},[]byte{2}))
	fmt.Println(bytes.Compare([]byte{},[]byte{}))

	//子切片sep在s中第一次出现的位置，不存在则返回-1。存在则返回序列数
	fmt.Println(bytes.Index([]byte("abcdef"),[]byte("d")))
	//判断切片b是否包含子切片subslice。
	fmt.Println(bytes.Contains([]byte("abcdef"),[]byte("p")))
}
