package main

import "fmt"

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

func main() {
	// go的函数中, 值类型的传递是直接复制, 引用类型(pointer, slice, map等)是指针复制, 内容不复制

	in := Data{
		complax: []int{1, 2, 3},
		instance: InnerData{
			5,
		},
		ptr: &InnerData{1},
	}

	fmt.Printf("in  value: %+v\n", in)
	fmt.Printf("in  ptr: %p\n", &in)

	out := passByValue(in)

	fmt.Printf("out  value: %+v\n", out)
	fmt.Printf("out  ptr: %p\n", &out)
}

func passByValue(inFunc Data) Data {
	// %+v 输出详细结构
	fmt.Printf("inFunc value: %+v\n", inFunc)
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	return inFunc
}
