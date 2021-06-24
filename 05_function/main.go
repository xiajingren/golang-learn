package main

import "fmt"

//普通函数
func fun1(a int) int {
	return a + 1
}

//多返回值函数
func fun2(a, b int) (int, int) {
	return b, a
}

//指定返回参数
func fun3() (a, b int) {
	a, b = 1, 2
	return
}

//声明函数类型
type myFuncType func(int) int

//函数作为参数
func fun4(fun myFuncType) int {
	return fun(100)
}

func main() {
	a := fun1(2)
	fmt.Println(a)

	c, d := fun2(5, 10)
	fmt.Print(c, d)

	e, f := fun3()
	fmt.Println(e, f)

	//匿名函数 函数变量
	fun := func(i int) int {
		fmt.Println("func4 callback")
		return i + 1
	}
	g := fun4(fun)
	fmt.Println(g)

	//立即执行
	func(s string) {
		fmt.Println(s)
	}("golang!")

	//闭包
	h := 10
	fun1 := func() {
		fmt.Println(h)
	}
	fun1()
	h = 11
	fun1()
}

/*
	3
	10 51 2
	func4 callback
	101
	golang!
	10
	11
*/
