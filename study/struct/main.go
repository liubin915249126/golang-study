package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 10
	cat1.Color = "白"
	cat1.Hobby = "吃"

	fmt.Println(cat1)
//创建结构体变量和访问结构体字段
//1.方式 1-直接声明
//案例演示: var person Person 前面我们已经说了。
//方式 2-{}

	p2 := Cat{"hhh", 20, "黑", "吃"}
	fmt.Println(p2)

	//方式 3-&
	var c3 *Cat = new(Cat)
	(*c3).Name = "dfsf"
	(*c3).Age = 10
	c3.Color = "红"
	(*c3).Hobby = "吃"

	fmt.Println(c3)

	//方式 4-{}
	var c4 *Cat = &Cat{}
	(*c4).Name = "dfsf"
	c4.Age = 100
	c4.Color = "红"
	c4.Hobby = "吃"

	fmt.Println(c4)

	//第 3 种和第 4 种方式返回的是 结构体指针。
	//结构体指针访问字段的标准方式应该是:(*结构体指针).字段名 ，比如 (*person).Name = "tom"
	//但 go 做了一个简化，也支持 结构体指针.字段名, 比如 person.Name = "tom"。更加符合程序员
	//使用的习惯，go 编译器底层 对 person.Name 做了转化 (*person).Name。
}