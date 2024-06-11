package main

import (
	"fmt"
)

func main() {
	var scores map [string] int // 声明了一个map变量，但是没有进行初始化，没有任何键值对
	//map值可以是任意类型，没有限制。map里所有键的数据类型必须是相同的，值也必须如何，但键和值的数据类型可以不相同。
	fmt.Printf("%T %#v\n",scores,scores)
    fmt.Println(scores == nil)  //表示是一个空值
	//映射格式 var m map[keyType]valueType


	//字面量  使用map字面量，需要先声明一个map变量
	//在Go语言中，map声明的字面量在本次使用之后是不能再被调用的。一旦字面量被使用，它的生命周期就结束了

	scores = map[string]int{"Cpu":8,"memory":9}
	fmt.Println(scores)


    /*可以使用make函数将值赋给变量，可以被多次调用，每次都会创建一个新的实例
    1表示map的初始容量而不是长度。初始容量是指在创建map时预留的存储空间，它可以提高map的性能。如果你知道大概会存储多少个键值对，
    可以通过设置初始容量来避免map频繁地重新分配内存。
	在本次案例，初始容量为1，表示在创建map时分配了1个存储桶。不过，这并不意味着map只能存储1个键值对，因为map会自动扩容以适应需要存储的
    键值对数量。即使初始容量为1，也可以存储多个键值队
     */
	scores = make(map[string]int,1)
	scores["apple"] = 1888
	scores["banana"] = 3
	scores["Gali"] = 20
	fmt.Println(scores,scores,scores)

	//map的增、删、改、查
	/*查询操作：可以使用value, ok := map[key]的方式来查询指定键的值。如果键存在，则将对应的值赋给value，并将ok设置为true；
	如果键不存在，则将value设置为该类型的零值，并将ok设置为false。
	一下有三种写法可以查询
	*/
	//使用value, ok := map[key]的方式来查询指定键的值，存在则执行if语句块并打印value。
	value,ok := scores["apple"]
	if ok {
		fmt.Println(value)
	}

	//-忽略返回的值。用于判断map中是否存在指定的键。如果键存在，则ok的值为true，否则为false。这种写法可以同时获取键对应的值和是否存在的信息。
	 _,okk := scores["Gal"]
		fmt.Println(okk)
	//判断map中是否存在指定的键，并将键对应的值赋给value变量。如果键存在，则进入if语句块并打印value的值。
	if value,ok := scores["banana"];ok{
		fmt.Println(value)
	}

	/* 增加（插入）操作：可以通过直接给map赋值或使用map[key] = value的方式来增加键值对。如果键已经存在，则会更新对应的值；
	如果键不存在，则会新增键值对。*/
	scores["Apiserver"]=18086
	fmt.Println(scores)



	//删除操作：可以使用delete(map, key)的方式来删除指定的键值对。如果键存在，则会删除对应的键值对；如果键不存在，则不做任何操作。
	//删除掉后，上面执行过的map打印的值不会有影响。只会在之后调用的map删除对应的键值队
	delete(scores,"Apiserver")
	fmt.Println(scores)


	//修改操作：通过给已存在的键赋新的值，即可修改对应的值。
	scores["apple"] = 6889
	fmt.Println(scores)

	//通过len获取指定map中的所有键值队的数量
	fmt.Println(len(scores))

	//通过键值队，使用for range循环打印map中的键值队。注意map是无序的
	for key,value := range scores {
		fmt.Println(key,value)
	}
//key至少可以有==，!= 运算，bool，整数，字符串，数组
//value = 任意类型 slice(切片) map(映射)
//名字==> 映射[字符串]{"地方、联系方式、成绩}

var user map[string]map[string]string //声明空map变量
	// map的类型 key key value，三种类型都是一样的。
	user = map[string]map[string]string{"吕良伟":{"地方":"山东","成绩":"86","联系方式":"123"}}
	fmt.Printf("%T,%#v\n,",user,user)

	user["李安南"] = map[string]string{"地方":"山西"}
	fmt.Println(user)

	user["李安南"] ["成绩"]= "99"
	fmt.Println(user)

	delete(user["吕良伟"],"联系方式")
	fmt.Println(user)

}

