package cls

import (
	"fmt"
	"xubo/chapter08/clsFactory/base"
)

type Class1 struct{}

func (c *Class1) Do() {
	fmt.Println("Class1")
}

func init() {
	base.Register("Class1", func() base.Class {
		return new(Class1)
	})
}
