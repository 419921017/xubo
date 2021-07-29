package session

import (
	"bufio"
	"net"
	"xubo/chapter12/12-6.tcppkt/packet"
)

func HandleSession(conn net.Conn, callback func(net.Conn, []byte) bool) {
	dataReader := bufio.NewReader(conn)
	for {
		pkt, err := packet.ReadPacket(dataReader)
		// 回调外部
		if err != nil || !callback(conn, pkt.Body) {
			conn.Close()
			break
		}
	}
}
