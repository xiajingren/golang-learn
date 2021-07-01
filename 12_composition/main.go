package main

import "fmt"

// type computer struct {
// 	model    string
// 	keyboard keyboard
// 	mouse    mouse
// }

type computer struct {
	model    string
	keyboard // 嵌入  (只写类型名 可以嵌入预申明类型int string ...)
	mouse    // 嵌入
}

type keyboard struct {
	name         string
	model        string
	NumberOfKeys int
}

type mouse struct {
	name  string
	model string
	dpi   int
}

func (keyboard) input() {
	fmt.Println("input...")
}

func (mouse) click() {
	fmt.Println("click...")
}

func (keyboard) info() {
	fmt.Println("keyboard info...")
}

func (mouse) info() {
	fmt.Println("mouse info...")
}

//转发到mouse的info方法
func (c computer) info() {
	c.mouse.info()
}

func main() {
	m := mouse{name: "mouse1", model: "m-001", dpi: 1000}
	fmt.Printf("%+v\n", m)
	k := keyboard{name: "keyboard1", model: "k-001", NumberOfKeys: 112}
	fmt.Printf("%+v\n", k)
	c := computer{model: "c-001", mouse: m, keyboard: k}
	fmt.Printf("%+v\n", c)

	// c.mouse.click()  //方法转发
	// c.keyboard.input()  //方法转发

	c.click()                  //通过嵌入 方法转发
	c.input()                  //通过嵌入 方法转发
	fmt.Printf("%+v\n", c.dpi) //也可以转发属性

	//当嵌入转发的方法(属性)出现重名时，自身的方法(属性)优先级高于转发的方法。
	//c.mouse.info()
	c.keyboard.info()
	// mouse 和 keyboard 都有info方法，如果使用 c.info 会发生歧义（不确定是调用mouse的还是keyboard的）
	//如果非要通过c.info访问mouse的info方法，那么给computer关联一个info方法，然后转发到mouse的info方法
	c.info()
}
