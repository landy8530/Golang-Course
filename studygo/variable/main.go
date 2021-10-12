package main

import (
	"fmt"
)

// 声明变量,建议使用驼峰标示
// var name string;
// var age int;
// var isOK bool;

// 批量声明变量
var (
	name string
	age  int
	isOK bool
)

func main() {

	name = "Landy"
	age = 18
	isOK = true

	//Go语言中声明了必须要使用，不使用就编译不过去

	fmt.Printf("name:%s\n", name)
	fmt.Printf("age:%d\n", age)
	fmt.Println("isOK", isOK)

	// 声明变量 同时 赋值

	var china string = "中国"
	fmt.Println("I love you :" + china)

	// 类型推导（根据值推断变量是什么类型）
	s := "aaaa"

	fmt.Println(s)

}
