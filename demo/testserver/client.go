package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil{
		fmt.Println("connect error")
	}
	_, err = conn.Write([]byte("just test"))
	if err != nil{
		fmt.Println("write error")
	}
}
