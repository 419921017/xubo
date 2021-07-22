package main

import (
	"fmt"
)

func main() {
	print(new(Alipay))
	print(new(Cash))
}

type Alipay struct{}

func (a *Alipay) CanUseFaceID() {}

type Cash struct{}

func (c *Cash) Stolen() {}

type ContainCanUseFaceID interface {
	CanUseFaceID()
}

type ContainStolen interface {
	Stolen()
}

func print(payMethod interface{}) {
	switch payMethod.(type) {
	case ContainCanUseFaceID:
		fmt.Printf("%T can use face id\n", payMethod)
	case ContainStolen:
		fmt.Printf("%T may be stolen", payMethod)
	}
}
