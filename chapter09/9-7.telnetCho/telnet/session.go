package telnet

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")

	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str)

			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))

		} else {
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}
