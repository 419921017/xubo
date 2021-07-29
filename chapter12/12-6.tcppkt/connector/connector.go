package connector

import (
	"fmt"
	"net"
	"strconv"
	"xubo/chapter12/12-6.tcppkt/packet"
)

func Connector(address string, sendTimes int) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < sendTimes; i++ {
		// 将数值转换成字符串
		str := strconv.Itoa(i)
		// 发送封包
		if err := packet.WritePacket(conn, []byte(str)); err != nil {
			fmt.Println(err)
			break
		}
	}
}
