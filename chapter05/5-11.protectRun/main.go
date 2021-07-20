package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("1-宕机前")
		panic(&PanicType{
			"触发panic",
		})
		fmt.Println("1-宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值-宕机前")
		var a *int
		*a = 1
		fmt.Println("赋值-宕机后")
	})
	fmt.Println("运行后")

}

type PanicType struct {
	function string
}

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error: ", err)
		default:
			fmt.Println("error: ", err)
		}

	}()
	entry()
}
