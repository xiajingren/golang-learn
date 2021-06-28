package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name string
	age  int
}

// 构造 person 的函数  构造函数命名格式通常为new开头
func newPerson(name string, age int) person {
	//to do something...

	return person{name, age}
}

// 为 person 关联方法
func (p person) say() {
	fmt.Println("my name is", p.name)
}

func main() {
	// 复合字面值
	p := person{name: "tom", age: 22}
	fmt.Println(p)

	var p1 person
	p1.name = "jack"
	p1.age = 23
	fmt.Println(p1)

	// struct 复制，将struct实例赋值给新的实例，修改新的实例不会影响原实例
	p2 := p1
	p2.age = 24
	fmt.Println(p1, p2)

	// struct 转换为json
	type book struct {
		Name, Author, explain string
	}
	b := book{Name: "book", Author: "someone", explain: "explain..."}
	// Marshal函数只会对struct中被导出的字段进行编码（首字母大写）
	bt, _ := json.Marshal(b)
	fmt.Println(string(bt))

	// 使用"标签"来解决json格式的字段命名问题
	type book1 struct {
		Name   string `json:"name"` //如果有多种格式 `json:"name"xml:"name"`
		Author string `json:"author"`
	}
	b1 := book1{Name: "book1", Author: "someone1"}
	bt1, _ := json.Marshal(b1)
	fmt.Println(string(bt1))

	// 构造函数
	p3 := newPerson("xh", 25)
	fmt.Println(p3)

	p3.say() // 调用关联方法
}
