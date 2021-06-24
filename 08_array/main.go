package main

import "fmt"

func main() {
	//定义数组
	a := [3]int{1, 2}
	fmt.Println(a)

	//使用...根据{}中的元素个数推数组导长度
	b := [...]int{1, 2, 3}
	fmt.Println(b)

	//将数组赋值给一个新数组，新数组中元素的改变不会影响原数组中元素的值
	c := [...]string{"hello", "word"}
	d := c
	d[1] = "golang"
	fmt.Println(c, d)

	//遍历数组
	for i := 0; i < len(c); i++ {
		fmt.Println(c[i])
	}
	//使用range遍历
	for i, v := range d {
		fmt.Println(i, v)
	}

	//二维数组
	e := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	for _, v1 := range e {
		for _, v2 := range v1 {
			fmt.Printf("%d ", v2)
		}
		fmt.Printf("\n")
	}
}
