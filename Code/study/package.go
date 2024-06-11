package main  //定义包

import "fmt" //导入包
//程序入口 main函数
//包级别
var packagevar string = "packagevar"
func main() {
	//包级别

		//函数级别的
		var funcvar string = "funcvar"
		var packagevar string = "func package var"
		fmt.Println("before",packagevar,funcvar)

			//块级别
			var blockvar string = "blockvar"
			{//限定变量的使用范围
			//子块级别定义的内部变量仅供子块使用，还可以使用外部定义的变量。外部使用不了子块(内部)定义的变量。go执行模块是由下而上
			var innerBlockvar string = "inner block var"
			fmt.Println(packagevar,funcvar,blockvar,innerBlockvar)
			}

			fmt.Println(packagevar,funcvar,blockvar)


		fmt.Println(funcvar, packagevar)

}