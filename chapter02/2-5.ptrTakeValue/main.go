package main

import "fmt"

func main() {
	var house = "Malibu Point 10800, 90265"

	ptr := &house

	// ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)
	// ptr的地址
	fmt.Printf("address: %p\n", ptr)
	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)
	// 指针取值
	fmt.Printf("value: %s\n", value)
}
