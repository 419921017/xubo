package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a *int
	fmt.Println("var a *int: ", reflect.ValueOf(a).IsNil())
	fmt.Println("nil: ", reflect.ValueOf(nil).IsValid())

	fmt.Println("(*int)(nil): ", reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct{}{}
	fmt.Println("不存在结构体成员: ", reflect.ValueOf(s).FieldByName("").IsValid())
	fmt.Println("不存在结构体方法: ", reflect.ValueOf(s).MethodByName("").IsValid())

	m := map[int]int{}

	fmt.Println("不存在的键: ", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}
