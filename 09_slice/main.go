package main

import "fmt"

func main() {
	//创建了一个长度为 3 的 int 数组，并返回一个切片给 a
	a := []int{1, 2, 3}
	fmt.Println(a, len(a), cap(a)) //cap:切片最大容量

	//从b[2]-b[4]创建切片
	b := []int{1, 2, 3, 4, 5}
	c := b[2:4]
	fmt.Println(c)
	//改变切片值  将会影响原数组
	c[0] = 22
	fmt.Println(b)

	//golang提供append函数来添加元素，当使用append函数时，
	//append函数会判断目的切片是否具有剩余空间，如果没有剩余空间，则会自动扩充两倍空间。
	d := []int{1, 2, 3, 4, 5}
	d = append(d, 6, 7, 8)
	fmt.Println(d)

	/*
		切片保留对底层数组的引用。只要切片存在于内存中，数组就不能被垃圾回收。这在内存管理方便可能是值得关注的。
		假设我们有一个非常大的数组，而我们只需要处理它的一小部分，为此我们创建这个数组的一个切片，并处理这个切片。
		这里要注意的事情是，数组仍然存在于内存中，因为切片正在引用它。

		解决该问题的一个方法是使用 copy 函数 func copy(dst, src []T) int 来创建该切片的一个拷贝。
		这样我们就可以使用这个新的切片，原来的数组可以被垃圾回收。
	*/
	e := [...]int{10, 11, 12, 13}
	f := e[:2]
	g := make([]int, len(f))
	copy(f, g) // 将f复制到g
	fmt.Println(e, f, g)

	//三索引切片
	h := [...]string{"aaa", "bbb", "ccc", "ddd", "eee"}
	i := h[1:3]                    //bbb ccc
	fmt.Println(i, len(i), cap(i)) // [bbb ccc] 2 4
	//开始索引，结束索引，容量限制索引
	j := h[1:3:3]
	fmt.Println(j, len(j), cap(j)) // [bbb ccc] 2 2
}
