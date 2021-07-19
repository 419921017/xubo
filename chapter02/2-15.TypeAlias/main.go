package main

import (
	"fmt"
	"reflect"
)

// 将NewInt定义为int类型
type NewInt int

// int取了个别名叫做IntAlias
type IntAlias = int

// type MyDuration = time.Duration

// func (m MyDuration) EasySet(a string) {
// 这里会报错, time.Duration和main包不是同一个包
// 将类型别名改成类型定义或者在time包下修改
// }

type Brand struct {
}

func (t Brand) Show() {

}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a NewInt
	fmt.Printf("a type: %T\n", a)
	var a2 IntAlias
	fmt.Printf("a2 type: %T\n", a2)

	var v Vehicle
	v.FakeBrand.Show()
	va := reflect.TypeOf(v)

	for i := 0; i < va.NumField(); i++ {
		f := va.Field(i)
		fmt.Printf("FileId Name: %v, FileId Type: %v\n", f.Name, f.Type.Name())
	}
}
