package main

import "fmt"

func main() {
	accumulator := Accumulate(1)
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	fmt.Printf("%p\n", accumulator)

	accumulator2 := Accumulate(10)
	fmt.Println(accumulator2())
	fmt.Println(accumulator2())

	fmt.Printf("%p\n", accumulator2)

}

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
