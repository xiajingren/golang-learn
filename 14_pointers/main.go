package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

func main() {
	answer := 42
	address := &answer
	fmt.Println(address, *address)
	fmt.Printf("%T\n", address)

	// 指针指向的变量发生变化，指针解引用的值也会跟着变化
	answer = 43
	fmt.Println(*address)

	// 修改指针解引用的值，原变量也会变化
	*address = 44
	fmt.Println(answer)

	// 两个指针指向同一个地址时，它们相互影响
	address1 := address
	*address1 = 45
	fmt.Println(*address, *address1)
	// 2个指针指向的地址相等，那么它们就是相等的
	fmt.Println(address == address1)

	// 指针解引用赋给另一个变量，相当于是复制了一份副本，此时两个变量互不影响
	answer1 := *address
	*address = 46
	fmt.Println(answer1, *address)

	timmy := &person{
		name: "Timothy",
		age:  10,
	}
	//(*timmy).superpower = "flying"
	// 访问字段时，对结构体指针进行解引用并不是必须的
	timmy.superpower = "flying" // 和上一行 等效
	fmt.Printf("%+v\n", timmy)

	// 指向数组的指针
	superpowers := &[3]string{"flight", "invisibilisy", "super strength"}
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca)

	//(&rebecca).birthday()
	rebecca.birthday() // 和上一行 等效
	fmt.Printf("%+v\n", rebecca)
}

func birthday(p *person) {
	p.age++
}

func (p *person) birthday() {
	p.age++
}
