package lesson

import "fmt"

func printA() {
	fmt.Println("a")
	fmt.Println("aa")
}

func printB() {
	fmt.Println("b")
}
func printC() {
	fmt.Println("c")
}

func DeferDemo() {
	// defer可以有多个，类似栈的逻辑，先压后出
	defer printA()
	defer printB()
	defer printC()
	fmt.Println("d")
	fmt.Println("e")
}

func deferFunc(a *int) *int {
	*a = *a + 1
	fmt.Printf("defer func called,a:%v\n", *a)
	return a
}

func returnFunc(a *int) *int {
	*a = *a * 2
	fmt.Printf("return func called,a:%v\n", *a)
	return a
}

// return func 先执行
func DeferDemo2() int {
	// 先执行return，再执行defer，因此return的结果是20
	a := 10
	defer deferFunc(&a)
	return *returnFunc(&a)
}
