package main

import (
	"net"
	"strconv"
	"xubo/chapter12/12-6.tcppkt/acceptor"
	"xubo/chapter12/12-6.tcppkt/connector"
)

func main() {
	const TestCount = 100000

	const address = "127.0.0.1:8010"

	var recvCounter int

	acceptor := acceptor.NewAcceptor()

	acceptor.OnSessionData = func(conn net.Conn, data []byte) bool {
		str := string(data)
		n, err := strconv.Atoi(str)

		if err != nil || recvCounter != n {
			panic("failed")
		}
		// 计数器增加
		recvCounter++
		// 完成计数, 关闭侦听器
		if recvCounter >= TestCount {
			acceptor.Stop()
			return false
		}
		return true
	}
	// 连接器不断发送数据
	connector.Connector(address, TestCount)
	// 等待侦听器结束
	acceptor.Wait()
}
