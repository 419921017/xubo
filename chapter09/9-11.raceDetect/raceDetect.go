package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var seq int64

func GenID() int64 {

	return atomic.AddInt64(&seq, 1)
}

var count int

var countGuard sync.Mutex

func GetCount() int {
	countGuard.Lock()
	defer countGuard.Unlock()
	return count
}
func SetCount(c int) {
	countGuard.Lock()
	defer countGuard.Unlock()
	count = c
}

func main() {
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID())

	SetCount(1)
	fmt.Println(GetCount())
}
