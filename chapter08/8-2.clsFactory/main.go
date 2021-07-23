package main

import (
	"xubo/chapter08/8-2.clsFactory/base"
	_ "xubo/chapter08/8-2.clsFactory/cls1"
	_ "xubo/chapter08/8-2.clsFactory/cls2"
)

func main() {

	c1 := base.Create("Class1")
	c1.Do()

	c2 := base.Create("Class1")
	c2.Do()
}
