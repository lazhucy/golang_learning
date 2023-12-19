package lesson

import "fmt"

const (
	BEIJING   int64 = 0
	SHANGHAI  int64 = 1
	CHONGQING int64 = 2
)

// 关键字iota,每行的iota都会累加1，第一行的iota默认为0
// iota只能配合const()使用
const (
	Luan    = 0
	Chuzhou = 1
	Anqin   = iota + 1
	Hefei
)

const (
	a1, b1 = iota + 1, iota + 2
	c1, d1
	e1, f1
	g1, h1 = iota * 2, iota * 3
	i1, j1
)

func Const() {
	fmt.Println(BEIJING, SHANGHAI, CHONGQING)
	fmt.Println(Luan, Chuzhou, Anqin, Hefei)
}
