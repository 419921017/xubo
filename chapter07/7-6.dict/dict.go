package main

import (
	"fmt"
)

func main() {
	dict := NewDictionary()
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite: ", favorite)
	dict.Visit(func(key, value interface{}) bool {
		if value.(int) > 40 {
			fmt.Println(key, "is expensive")
			return true
		}
		fmt.Println(key, "is cheap")
		return true
	})
}

type Dictionary struct {
	data map[interface{}]interface{}
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}
