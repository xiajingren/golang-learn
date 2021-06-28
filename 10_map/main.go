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
	// 用 ok 来判断 key 是否存在
	if m, ok := a["wuhan"]; ok {
		fmt.Println(m)
	} else {
		fmt.Println("wuhan not found...")
	}

	// 删除map中的key
	delete(a, "guangzhou")
	fmt.Println(a)
}
