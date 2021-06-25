# golang 基础



## hello word

创建项目，使用**go mod** （*golang推荐的包管理工具）

`go mod init 01_hello_golang` //初始化mod

新建main.go

```go
package main //包名 main

import "fmt" //导入 fmt 包

//入口函数 main ， main函数必须定义在 main 包中
func main() {
	fmt.Println("hello golang !") //输出文本 换行结尾
}
```

`go run main.go` //编译运行

`go build main.go` //打包为可执行文件



## 变量 var

```go
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
```



## 基础数据类型 data type

- bool

- int

  *u 有符号/无符号 uint8,uint16,uint32,uint64,int8,int16,int32,int64

- float32/float64

- string

- byte(uint8)

- rune(int32)

- ......



## 格式化输出动词 format_output_verb

%v	按值本身的格式输出，万能动词，不知道用什么动词就用它
%+v	同%v，但输出结构体时会添加字段名
%#v	输出值得Go语法表示
%t	格式化bool值
%b	按二进制输出
%c	输出整数对应的Unicode字符
%d	按十进制输出
%o	按八进制输出
%O	按八进制输出，并添加0o前缀
%x	按十六进制输出，字母小写，%x还能用来格式化字符串和[]byte类型，每个字节用两位十六进制表示，字母小写
%X	按十六进制输出，字母大写，%X还能用来格式化字符串和[]byte类型，每个字节用两位十六进制表示，字母小写
%U	按Unicode格式输出
%e	按科学计数法表示输出，字母小写
%E	按科学计数法表示输出，字母大写
%f	输出浮点数
%F	同%f
%g	漂亮的格式化浮点数
%G	同%G
%s	格式化为字符串
%q	格式化为字符串，并在两端添加双引号
%p	格式化指针
%T	输出变量类型
%w	专用于fmt.Errorf包装error



## 常量 const

```go
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
```



## 函数 function

- 头等函数

  > 支持头等函数的语言允许将函数赋值给变量，作为参数传递给其他函数并从其他函数返回。 Go 支持头等函数功能。

- 匿名函数

  > 匿名函数就是没有名字的函数，在Go里也称作函数字面值。
  >
  > 因为函数字面值需要保留外部作用域的变量引用，所以函数字面值都是闭包的。

- 闭包

  > 闭包（closure）就是由于匿名函数封闭并包围作用域中的变量而得名的。

  

```go
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
```



## 方法 method

在其他面向对象语言里，方法属于类。

在Go里，提供了方法，但是没有提供类和对象，Go中的方法是与类型进行关联。

```go
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
```



## defer

```go
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
```



## 数组 array

- 数组是类型相同的元素的集合，Go不允许在数组中混合使用不同类型的元素（比如整数和字符串）。
- 数组一旦声明，长度不可改变，数组的长度是数组类型的一部分。

- 在 Go 中数组是值类型而不是引用类型。这意味着当数组变量被赋值时，将会获得原数组的拷贝。新数组中元素的改变不会影响原数组中元素的值。

```go
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
```



## 切片 silce

> 切片（slice）是建立在数组之上的更方便，更灵活，更强大的数据结构（窗口，视图）。切片并不存储任何元素而只是对现有数组的引用。



golang中数组和切片的区别：

●　切片是指针类型，数组是值类型

●　数组的长度是固定的，而切片不是（切片是动态的数组）

●　切片比数组多一个属性：容量（cap)

●　切片的底层是数组

```go
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
```



## map

```go
package main

import "fmt"

// map 作为参数时 是引用传递。在其他方法中修改 会影响原map (赋值同理)
func fun1(m map[string]string) {
	m["guangzhou"] = "gz"
}

func main() {
	// 定义map
	a := map[string]string{
		"shanghai": "sh",
		"beijing":  "bj",
		"shenzhen": "sz",
	}
	fmt.Println(a)
	fun1(a)
	fmt.Println(a)

	// 访问不存在的key，value会返回该数据类型的默认值
	fmt.Println("wuhan=", a["wh"])

	if m, ok := a["wuhan"]; ok {
		fmt.Println(m)
	} else {
		fmt.Println("wuhan not found...")
	}

	// 删除map中的key
	delete(a, "guangzhou")
	fmt.Println(a)
}
```



