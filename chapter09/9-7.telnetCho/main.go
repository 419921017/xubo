package main

import (
	"fmt"
	"os"
	"xubo/chapter09/9-7.telnetCho/telnet"
)

func main() {
	exitChan := make(chan int)
	telnet.Server("127.0.0.1:7001", exitChan)
	code := <-exitChan
	fmt.Println("code", code)
	os.Exit(code)
}
