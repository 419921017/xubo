package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	// 使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。
	// Range参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回true；终止迭代遍历时，返回false
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})
}
