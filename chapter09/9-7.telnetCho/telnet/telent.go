package telnet

import (
	"fmt"
	"strings"
)

func processTelnetCommand(str string, exitChan chan int) bool {
	// @close表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
		// @shutdown表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		exitChan <- 0
		return false
	}

	fmt.Println(str)
	return true

}
