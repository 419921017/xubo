package main

import (
	"time"

	"github.com/pkg/profile"
)

func main() {
	stropper := profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer stropper.Stop()

	joinSlice()
	time.Sleep(time.Second)
}

func joinSlice() []string {
	var arr []string
	for i := 0; i < 100000; i++ {
		arr = append(arr, "arr")
	}
	return arr
}
