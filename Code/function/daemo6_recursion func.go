package main

import "fmt"
//汉诺塔游戏。源A 借助B ——> 目的C
func tower(a,b,c string,layer int)  {
	if layer == 1{
		fmt.Println(a,"->",c)
		return
	}

	tower(a, c ,b, layer-1)
	fmt.Println(a,"->",c)
	tower(b,a,c, layer-1)

}
func main() {
	fmt.Println("1层:")
	tower("A","B","C",1)
	fmt.Println("2层:")
	tower("A","B","C",2)
	fmt.Println("3层:")
	tower("A","B","C",3)
	fmt.Println("4层:")
	tower("A","B","C",4)
	fmt.Println("5层:")
	tower("A","B","C",5)
}
