package main

import (
	"fmt"
)

// 定义一个接口变量
var t interface {
	talk() string
}

// 定义结构体
type imp1 struct{}
type imp2 struct{}

//为imp1关联talk方法
func (imp1) talk() string {
	return "imp1"
}

//为imp2关联talk方法
func (imp2) talk() string {
	return "imp2"
}

// 声明接口类型
type talker interface {
	talk() string
}

func shout(talk talker) {
	fmt.Println(talk.talk())
}

type starship struct {
	imp1 // 嵌入 imp1 类型
}

func main() {
	t = imp1{} //imp1满足t这个接口
	fmt.Println(t.talk())
	t = imp2{} //imp2满足t这个接口
	fmt.Println(t.talk())

	imp1 := imp1{} //imp1满足talker这个接口
	shout(imp1)
	imp2 := imp2{} //imp2满足talker这个接口
	shout(imp2)

	// 因为 starship 嵌入了 imp1 类型，所以转发了 talk 方法
	s := starship{}
	fmt.Println(s.talk()) //所以 starship 可以直接调用 talk 方法
	shout(s)              //所以 starship 满足了 talker 这个接口，所以可以把 s 直接传给 shout 方法
}
