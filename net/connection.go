package net

import (
	"fmt"
	"net"
	"server/iface"
)

type Conection struct {
	Conn *net.TCPConn
	ConnId uint32
	IsClosed bool
	HandleFunc iface.HandleFunc
	ExitChan chan bool
}

func WorkConn(c *Conection){
	if c.IsClosed{
		return
	}
	defer c.Conn.Close()
	for {
		buf := make([]byte,512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read failed .....")
			continue
		}
		c.HandleFunc(c.Conn,buf,cnt)
	}

}

func (c *Conection) Start()  {
	fmt.Println("conn star ....,conn id = %d",c.ConnId)
	go WorkConn(c)
}

func (c *Conection) Stop()  {
	if c.IsClosed {
		return
	}
	c.Conn.Close()
	c.IsClosed = true
}

func (c *Conection) GetTcpConnection() *net.TCPConn{
	return c.Conn
}

func (c *Conection) GetConnId() uint32{
	return c.ConnId
}

func (c *Conection) RemoteAddr() net.Addr{
	return c.Conn.RemoteAddr()
}

func NewConnection(conn *net.TCPConn, connID uint32, call_api iface.HandleFunc) *Conection{
	c := &Conection{
		Conn: conn,
		ConnId: connID,
		HandleFunc: call_api,
		IsClosed: false,
		ExitChan: make(chan bool),
	}
	return c
}
