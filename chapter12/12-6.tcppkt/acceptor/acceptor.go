package acceptor

import (
	"fmt"
	"net"
	"sync"
	"xubo/chapter12/12-6.tcppkt/session"
)

type Acceptor struct {
	l             net.Listener
	wg            sync.WaitGroup
	OnSessionData func(net.Conn, []byte) bool
}

func (a *Acceptor) Start(address string) {
	go a.listen(address)
}

func (a *Acceptor) listen(address string) {
	a.wg.Add(1)
	defer a.wg.Done()

	var err error

	a.l, err = net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		conn, err := a.l.Accept()

		if err != nil {
			break
		}
		// 根据连接发起会话
		go session.HandleSession(conn, a.OnSessionData)
	}

}

func (a *Acceptor) Stop() {
	a.l.Close()
}

func (a *Acceptor) Wait() {
	a.wg.Wait()
}

func NewAcceptor() *Acceptor {
	return &Acceptor{}
}
