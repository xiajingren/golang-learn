package main

import "fmt"

var z = 999 //全局变量

func main() {
	//声明变量名称，类型 不赋值的话默认值为该数据类型的“零值”
	var a int
	fmt.Println(a)

	//声明变量名称，类型，并初始化变量值
	var b int = 5
	fmt.Println(b)

	//自动推导变量类型
	var c = 6
	fmt.Println(c)

	//声明多个变量
	var d, e = 2, 3
	var (
		f = 4
		g = "str"
	)
	fmt.Println(d, e, f, g)

	//短声明（常用） *但是不支持全局变量的声明
	h := 10
	i := 11
	fmt.Println(h, i)

	fmt.Println(z)
}
