package main

import "fmt"

func main() {
	print(1, 2, 3)
}

func print(slist ...interface{}) {
	rawPrint(slist...)
}

func rawPrint(rawList ...interface{}) {
	for _, v := range rawList {
		fmt.Println(v)
	}
}
