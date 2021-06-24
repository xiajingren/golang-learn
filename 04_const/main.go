package main

import "fmt"

func main() {
	//定义常量
	const a = 1
	fmt.Println(a)

	//定义多个常量
	const (
		bj = 10
		sh //不赋值 默认和上一行相等
		sz //不赋值 默认和上一行相等
	)
	//10 10 10
	fmt.Println(bj, sh, sz)

	//iota 每当定义一个const，iota的初始值为0，逐行累加，可以参与计算
	const (
		aa = iota * 10
		bb
		cc
	)
	//0 10 20
	fmt.Println(aa, bb, cc)

	//iota 逐行累加，不是根据常量累加
	const (
		dd, ee = iota + 1, iota + 2
		ff, gg
	)
	//1 2 2 3
	fmt.Println(dd, ee, ff, gg)
}
