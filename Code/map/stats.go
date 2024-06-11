package main

import "fmt"

func main() {
	user := []string {"庄肃磊","翔哥","吕良伟","鸡哥","吕良伟","庄肃磊","翔哥","庄肃磊",}

	voteCount := make(map[string]int)
	for _, name := range user {
		voteCount[name]++
	}

	maxVotes := 0
	winner := ""
	for name, votes := range voteCount {
		if votes > maxVotes {
			maxVotes = votes
			winner = name
		}
	}
	fmt.Println("得票最高的人是：", winner)
	/*
	这段代码的作用是统计每个人的得票数，并找出得票最高的人。

	voteCount := make(map[string]int)：创建一个空的map，用于存储每个人的得票数。这里的键是人的名字，值是得票数。

	for _, name := range user：遍历名字切片user，对于每个名字，执行以下操作：

	voteCount[name]++：通过键name访问mapvoteCount，将对应的值加1。如果该名字在map中不存在，会被自动初始化为0，然后再加1。
	maxVotes := 0：初始化最大得票数为0。

	winner := ""：初始化得票最高的人为空字符串。

	for name, votes := range voteCount：遍历voteCount中的键值对，对于每个键值对，执行以下操作：

	if votes > maxVotes：判断当前得票数是否大于最大得票数。

	如果是，则将当前得票数赋给最大得票数maxVotes，将对应的人名赋给得票最高的人winner。

	如果不是，则继续遍历下一个键值对。

	fmt.Println("得票最高的人是：", winner)：打印得票最高的人。

	这段代码通过遍历名字切片，并使用map统计每个人的得票数。然后，通过比较得票数找出最高得票的人。最后，打印得票最高的人。
	 */


}
