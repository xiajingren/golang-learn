package main

import "fmt"

// defer在return之后执行 当defer被声明时，其参数就会被实时解析
func a() {
	i := 0
	defer fmt.Println(i) // 0
	i++
	//return
}

// defer先定义后执行
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Println(i) // 3 2 1 0
	}
}

// defer可以读取带名称返回值
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	b()
	fmt.Println(c()) // 2  (闭包 ???)
}
