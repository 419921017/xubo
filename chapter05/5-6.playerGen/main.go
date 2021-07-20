package main

import "fmt"

func main() {
	generator := playerGen("high noon")
	name, hp := generator()

	fmt.Println(name, hp)
}

func playerGen(name string) func() (string, int) {
	hp := 150

	return func() (string, int) {
		return name, hp
	}
}
