package main

import "fmt"

func main() {
	var progress = 2
	var target = 8

	title := fmt.Sprintf("%d %d\n", progress, target)
	fmt.Println(title)

	pi := 3.14159

	variant := fmt.Sprintf("%v, %v, %v\n", "moon", pi, true)

	fmt.Println(variant)

	profile := &struct {
		Name string
		Hp   int
	}{
		Name: "rat",
		Hp:   150,
	}

	// %+v, 对结构体字段名和值进行展开
	fmt.Printf("'%%+v' %+v\n", profile)
	// 输出语法格式的值
	fmt.Printf("'%%#v' %#v\n", profile)
	//
	fmt.Printf("'%%T' %T\n", profile)

}
