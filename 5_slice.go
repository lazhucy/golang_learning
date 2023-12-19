package lesson

import "fmt"

func Slices() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("len:%d,slice1:%v\n", len(slice1), slice1)

	// 声明slice的类型，但是没有分配空间
	var slice2 []int
	// 通过make给slice开辟空间
	slice2 = make([]int, 0, 3)
	fmt.Printf("len:%d,slice2:%v\n", len(slice2), slice2)

	var slice3 []int = make([]int, 0, 3)
	fmt.Printf("len:%d,slice3:%v\n", len(slice3), slice3)

	slice4 := make([]int64, 0, 3)
	fmt.Printf("len:%d,slice4:%v\n", len(slice4), slice4)

	slice5 := make([]int64, 3, 5)
	//slice5[3] = 10 // 只是预分配了空间，但是数组长度只有3，这种操作数组越界
	slice5 = append(slice5, 10, 11)
	fmt.Printf("len:%d,cap:%d,slice5:%v\n", len(slice4), cap(slice5), slice5)
	slice5 = append(slice5, 12, 13)
	fmt.Printf("len:%d,cap:%d,slice5:%v\n", len(slice4), cap(slice5), slice5)
	s1 := slice5[0:2]
	s2 := slice5[:3]
	s3 := slice5[2:]
	fmt.Println(s1, s2, s3)
	// 浅拷贝,s,s1,s2首个元素都被改成100了，这个要注意！
	s1[0] = 100
	fmt.Println(slice5, s1, s2)
	// 深拷贝,将s1进行值传递
	s4 := make([]int64, 3)
	copy(s4, s1) // s1中的数依次拷贝进s4中，s4的长度小于s1，则拷到长度上限为止，s4不会自动扩容
	fmt.Println(s1, s4)
	s4[0] = 101
	fmt.Println(slice5, s1, s2, s4)
}
