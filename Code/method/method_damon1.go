package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) hello() {
	fmt.Printf("hello ! my name is %v\n", p.Name)
}

func (p Person) speak() {
	fmt.Printf("%v 是一个好人！\n", p.Name)
}

func (p Person) calc() {
	sum := 0
	for i := 0; i < 1000; i++ {
		sum += i
	}
	fmt.Printf("the result id %d", sum)
}

//func (p Person) calc1(n int) {
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += i
//	}
//	fmt.Printf("the result id %d", sum)
//}

func (p Person) getSum(a int, b int) int {
	return a + b
}

func main() {
	var p Person
	p.Name = "jerry"
	p.Age = 55
	p.hello()
	p.speak()
	p.calc()
	p.getSum(1, 1)
	fmt.Println(p.getSum)

	person := Person{Name: "john", Age: 30}
	sum := p.getSum(20, 3)
	fmt.Printf("\n打印id", person)
	fmt.Printf("\n结果为", sum)

}
