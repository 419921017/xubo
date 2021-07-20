package main

import (
	"fmt"
)

func main() {
OuterLoop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				// break OuterLoop
				continue OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
