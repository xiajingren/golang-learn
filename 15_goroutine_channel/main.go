package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 创建通道
	c := make(chan int)

	// 开启 5 个 goroutine 、 使用 go 关键字
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c) // 开启一个 goroutine 不会阻塞
	}

	// 从通道接收 5 次
	for i := 0; i < 5; i++ {
		gopherId := <-c // 从通道 c 接收数据，会阻塞。可以用于判断 goroutine 是否执行完毕（有点像 await）
		fmt.Println("gopher ", gopherId, " has finished sleeping")
	}

	fmt.Println("=======================================")

	c1 := make(chan int)
	// time.After 返回一个通道，该通道在指定时间后会接收到一个值，发送该值的 goroutine 是 go 运行时的一部分
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		go sleepyGopher1(i, c1)
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c1: // 从 c1 通道等待一个值
			fmt.Println("gopher ", gopherId, " has finished sleeping")
		case <-timeout: // c1 通道 2s 没有值则会走此 case
			fmt.Println("my patience ran out")
			//return
		}
	}

	fmt.Println("=======================================")
	c2 := make(chan int)
	go sleepyGopher2(c2)
	// 使用 range 关键字 从通道接收数据，直到通道被关闭
	for a := range c2 {
		fmt.Println(a)
	}

	fmt.Println("end...")
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second) // 休眠 3s
	fmt.Println("...", id, " snore ...")
	c <- id // 将 id 发送到通道 c 代表当前方法执行完毕
}

func sleepyGopher1(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // 随机休眠
	fmt.Println("...", id, " snore ...")
	c <- id // 将 id 发送到通道 c 代表当前方法执行完毕
}

func sleepyGopher2(c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // 随机休眠
		fmt.Println("...", i, " snore ...")
		c <- i // 将 id 发送到通道 c 代表当前方法执行完毕
	}
	close(c)
}
