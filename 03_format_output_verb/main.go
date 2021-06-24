package main

import (
	"errors"
	"fmt"
)

type point struct {
	x, y, z int
}

var (
	p = point{1, 2, 3}
)

func main() {
	//{1 2 3}
	fmt.Printf("%v\n", p)

	//{x:1 y:2 z:3}
	fmt.Printf("%+v\n", p)

	//main.point{x:1, y:2, z:3}
	fmt.Printf("%#v\n", p)

	//true
	fmt.Printf("%t\n", true)

	//1010
	fmt.Printf("%b\n", 10)

	//A
	fmt.Printf("%c\n", 65)

	//1
	fmt.Printf("%d\n", 1)

	//11
	fmt.Printf("%o\n", 9)

	//0o11
	fmt.Printf("%O\n", 9)

	//a
	fmt.Printf("%x\n", 10)

	//A
	fmt.Printf("%X\n", 10)

	//010a
	fmt.Printf("%x\n", []byte{1, 10})

	//010A
	fmt.Printf("%X\n", []byte{1, 10})

	//68656c6c6f
	fmt.Printf("%x\n", "hello")

	//776F726C64
	fmt.Printf("%X\n", "world")

	//U+0001
	fmt.Printf("%U\n", 1)

	//1.200000e+00
	fmt.Printf("%e\n", 1.2)

	//1.200000E+00
	fmt.Printf("%E\n", 1.2)

	//1.200000
	fmt.Printf("%f\n", 1.2)

	//1.200000
	fmt.Printf("%F\n", 1.2)

	//1.2
	fmt.Printf("%g\n", 1.2)

	//1.2
	fmt.Printf("%G\n", 1.2)

	//hello
	fmt.Printf("%s\n", "hello")

	//"world"
	fmt.Printf("%q\n", "world")

	//0x563820
	fmt.Printf("%p\n", &p)

	//main.point
	fmt.Printf("%T\n", p)

	//outer:inner
	fmt.Println(fmt.Errorf("outer:%w", errors.New("inner")))
}
