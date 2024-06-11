package main //定义包,包名在被跨目录调用时要和目录名保持一致。一个目录下有多个Go文件，第二个文件不能和第一个文件不能同时命名main包，建议第二个文件和目录名称保持一致

import (
	. "fmt"
) //导入标准包

var Lvliangwei string = "test"

// 程序入口 main函数
func main() {
	//变量.给数据起一个名字
	var msg string = "hello-world"
	/*这是我的第一个程序
	这里调用fmt包下的println函数打印内容到控制台
	*/
	Println(msg) //打印helo-word内容到控制台
	//fmt.Println(msg) //打印helo-word内容到控制台
	//fmt.Println(msg) //打印helo-word内容到控制台
	//fmt.Println(msg) //打印helo-word内容到控制台
	//fmt.Println(msg) //打印helo-word内容到控制台
	//fmt.Println(msg) //打印helo-word内容到控制台

}

/*点操作
import(
. "fmt"
)
这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调

用的`fmt.Println("hello world")`可以省略的写成`Println("hello world")`
*/
