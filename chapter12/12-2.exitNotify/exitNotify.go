package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func socketRecv(conn net.Conn, wg *sync.WaitGroup) {
	buff := make([]byte, 1024)
	for {
		_, err := conn.Read(buff)
		if err != nil {
			break
		}
	}

	wg.Done()
}

func main() {
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go socketRecv(conn, &wg)

	time.Sleep(time.Second)

	conn.Close()

	wg.Wait()

	fmt.Println("recv done")
}
