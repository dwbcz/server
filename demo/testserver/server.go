package main

import (
	"server/net"
)

func main()  {
	s := net.NewServer("test")
	s.Serve()
}