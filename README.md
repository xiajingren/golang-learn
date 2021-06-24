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











