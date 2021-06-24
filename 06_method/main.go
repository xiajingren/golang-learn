package main

import "fmt"

//定义person结构体（类型）
type person struct {
	name string
	age  int
}

//为person类型关联方法 p:类型参数接收者，方法可以有多个参数返回值，但只能有一个接收者
//不能为int float等预声明的的类型进行关联，(起别名后再进行关联)
func (p person) say() {
	fmt.Printf("%+v\n", p)
}

func main() {
	p := person{name: "jack", age: 22}
	p.say() //调用
}
