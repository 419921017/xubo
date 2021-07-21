package main

import "fmt"

type Wheel struct {
	Size int
}

// type Engine struct {
// 	Power int
// 	Type  string
// }
type Car struct {
	Wheel
	Engine struct {
		Power int
		Type  string
	}
}

func main() {
	c := Car{
		Wheel: Wheel{
			Size: 18,
		},
		Engine: struct {
			Power int
			Type  string
		}{
			Power: 153,
			Type:  "1.4T",
		},
	}

	fmt.Printf("%+v", c)
}
