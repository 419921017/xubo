package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
}

type file struct{}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("Write Data: ", data)
	return nil
}

func main() {
	f := new(file)

	var writer DataWriter = f

	writer.WriteData("data")
}
