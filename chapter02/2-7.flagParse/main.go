package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()
	fmt.Println(*mode)

	// new 会直接创建一个对应的指针的类型
	str := new(string)
	fmt.Printf("str type:  %T\n", str)
	*str = "ninja"
	fmt.Println(*str)
}
