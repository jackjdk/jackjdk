package main

import "fmt"

func main() {
	var a1 A = Cat{"狸花猫"}
	var a2 A = Person{"王二狗", 30}
	var a3 A = "haha"
	var a4 A = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

}

type A interface {
}

type Cat struct {
	color string
}
type Person struct {
	name string
	age  int
}
