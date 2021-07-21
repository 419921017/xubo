package main

import "fmt"

type Wheel struct {
	Size int
}
type Engine struct {
	Power int
	Type  string
}
type Car struct {
	Wheel
	Engine
}

// 初始化结构体内嵌
func main() {
	c := Car{
		Wheel: Wheel{
			Size: 18,
		},
		Engine: Engine{
			Type:  "1.4T",
			Power: 143,
		},
	}

	fmt.Printf("%+v\n", c)
}
