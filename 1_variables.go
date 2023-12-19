package lesson

import "fmt"

var e float32 = 0.2

func Variables() {
	// 四种初始化方法
	// 前三种可以用来申明全局变量
	var a int
	var b int = 32
	var c = 32
	d := 100 // 最常用

	// 多变量声明
	var (
		h, i int64 = 12, 23
		j    bool  = false
	)
	fmt.Printf("type of a:%T\ntype of b:%T\ntype of c:%T\ntype of d:%T\n", a, b, c, d)
	fmt.Println(e)
	fmt.Println(h, i, j)
	fmt.Println("Hello World!")
}
