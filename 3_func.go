package lesson

import "fmt"

// 多个返回值，匿名的
func Foo1(a string, b int) (string, int) {
	c := a + "hello"
	d := b * 2
	return c, d
}

// 多个返回，有形参名的，不赋值就是默认值
func Foo2(a string, b int) (c string, d int) {
	c = a + "hello"
	d = b * 2
	return
}

func Foo3(a int) (res map[int]int) {
	fmt.Println(res == nil)
	fmt.Println(res[0])
	var tmp map[int]int
	fmt.Println(tmp == nil)
	tmp = make(map[int]int, 0)
	fmt.Println(tmp == nil)
	//fmt.Println("res equal to tmp:",res = tmp)
	fmt.Println(tmp[1])
	return
}
