package net

import (
	"fmt"
	"net"
	"server/iface"
)

type Server struct {
	IP string
	Port int
	IPVersion string
	Name string
}

func (s *Server) Serve()  {
	s.Start()

	//TODO to some work
	select {}
	s.Stop()
}

func (s *Server) Start()  {
	fmt.Printf("server starting ...， IP:%s, Port: %d",s.IP,s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err != nil {
			fmt.Println("resolve IP addr error ....")
		}

		//Listen
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp faild .....")
		}
		fmt.Println("start server success, name = %s .....",s.Name)

		//建立链接
		for {
			cli, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("conne cli error ....")
				continue
			}
			go func() {
				for {
					buf := make([]byte, 1024)
					if _, err := cli.Read(buf); err != nil {
						fmt.Println("read message error")
					} else {
						wlen, err := cli.Write(buf)
						if err != nil {
							fmt.Println("write failed .....")
						}
						fmt.Println("mesaage : %s , len = %d", buf, wlen)
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop()  {
	//TODO stop server
}

func NewServer(name string) iface.IServer{
	s := &Server{
		IP: "127.0.0.1",
		Port: 8080,
		IPVersion: "ip4",
		Name: name }
	return s
}

