package main

import "fmt"

const (
	OK        = 200
	NOT_FOUND = 404
)

// 批量声明常量时，如果某一行没有赋值，则默认为上一行的值
// const (
// 	n1 = 100
// 	n2
// 	n3
// )
// iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	n1 = iota //0
	n2        //1
	n3        //2
	n4        //3
)

func main() {

	// 定义常量
	// 定义之后不能修改
	const pi = 3.1415926

	fmt.Println(pi)

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
}
