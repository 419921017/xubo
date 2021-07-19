package main

import "fmt"

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}
