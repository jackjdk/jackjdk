package main

import "fmt"

func main() {

/*
	字符串:
	1.概念:多个byte的集合。理解为一个字符序列
	2.语法:使用双引
	"abc","hello","A"
	也可以使用``
	3.编码问题
			计算机本质只识别0和1
			A：65，B：66，C：67
			a:97，b：98
			ASCII(美国标准信息交换码)
			中国的编码表：gpk，兼容ASCII
			中，家
			Unicode编码表：号称统一了全世界
				UTF-8，UTF-16,UTF-32..
	4.转义字符：\
		A:有一些字符，有特殊的作用，可以转义为普通的字符
			\',\'
		B:有一些字符，就是一个普通的字符，转义后有特殊的作用
			\n，换行
			\t，制表符
 */
//1.定义字符串
var s1 string = "王二狗"
fmt.Println(s1)

//2.区别："A","A"
v1 :='A'
v2 := "A"

fmt.Printf("%T,%d\n",v1,v1)
fmt.Printf("%T,%s\n",v2,v2)

//3.转义符:\

fmt.Println("\"hello word\"")
fmt.Println("hello\nwor\td")

fmt.Println(`He"llo wor"d`)
fmt.Println("hello`word`rd")






	/*
	//int:比较字符串 a 和字符串 b，在 a<b 时返回-1，在a>b 时返回1，在a=b 时返回0。
	fmt.Println(strings.Compare("a","b"))
	//bool：判断字符串 s 中是否包含子串 substr，返回布尔值
	fmt.Println(strings.Contains("s","substr"))
	//bool:判断字符串 s 中是否包含子串 chars, 中的其中一个字符，返回布尔值
	fmt.Println(strings.ContainsAny("c","abc"))
	//int:统计字符串 s 中子串 substr 出现的次数。
	fmt.Println(strings.Count("sbasba","sbasba"))
	//string:将字符串s按照空格分割成多个字符串并返回结果（类似还有 FieldsFunc）。
	fmt.Println(strings.Fields("abc defn\neee\raaa\fxxx\vsdd"))
	fmt.Printf("%q\n",strings.Fields("abc def\neee\raaaa\fxxx\vsdd"))

	\n：换行符（newline），用于在输出中换行。
	\r：回车符（carriage return），用于将光标移到行首。
	\f：换页符（form feed），用于在输出中进行分页。
	\v：纵向制表符（vertical tab），用于在输出中进行垂直制表。

	//bool:判断字符串 s 是否以 prefix 开头，返回布尔值。
	fmt.Println(strings.HasPrefix("pre","prefix"))
	//bool:判断字符串 s 是否以 suffix 结尾，返回布尔值
	fmt.Println(strings.HasSuffix("x","suffix"))
	//int：返回字符串 s 中子串 substr 第一次出现的位置，如果没有找到则返回-1
	fmt.Println(strings.Index("subst","subst"))
	//int：返回字符串 s 中子串 substr 最后一次出现的位置，如果没有找到则返回-1
	fmt.Println(strings.LastIndex("acdtssssacdtacdt","acdt"))
	//string：将字符串 s 按照 sep 分割成字符串切片（类似还有 SplitN）。
	fmt.Println(strings.Split("abcdef;abc;absss","ab"))
	//将一个字符串切片中的所有元素用指定的分隔符连接起来
	fmt.Println(strings.Join([]string{"abc","def"},""))

	fmt.Println(strings.Repeat("abc",4))
	//string：将字符串 s 中的前 n 个 old 替换为 new，n<0 时则替换所有 old 为 new。
	fmt.Println(strings.Replace("abcdabcdabcd","ab","xx",2))
	//-1替换全部
	fmt.Println(strings.Replace("abcdabcdabcd","ab","xx",-1))
	//string：将字符串 s 中的所有 old 替换为 new
	fmt.Println(strings.ReplaceAll("abcdabcdabcd","ab","xx",))
	//string：将字符串 s 中的所有字母转换为小写
	fmt.Println(strings.ToLower("aaaa"))
	//string：将字符串 s 中的所有字母转换为大写
	fmt.Println(strings.ToUpper("aaaa"))
	//string：将字符串 s 中的每个单词的首字母转换为大写
	fmt.Println(strings.Title("hi,ok"))
	//string：去除字符串 s 两侧的 cutset 字符集中的字符（类似的还有 TrimLeft
	fmt.Println(strings.Trim("abcdefz","az"))
	//去除字符串 s 两侧的空格符（或其他 unicode.IsSpace 返回 true 的字符）。
	fmt.Println(strings.TrimSpace("abcd xxx \n aaa \r z zzzz "))

*/












}
