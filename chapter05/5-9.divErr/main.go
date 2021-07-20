package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(div(1, 0))
}

var errDivisionByZero = errors.New("division by zero")

func div(divided, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errDivisionByZero
	}

	return divided / divisor, nil
}
