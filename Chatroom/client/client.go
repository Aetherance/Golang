package main

import "net"

type Client struct {
	addr string
}

func NewClient() * Client {
	return &Client{
		
	}
}

func (c * Client) setAddr(addr string) {
	c.addr = addr
}

func (c * Client) Connect() {
	if len(c.addr) == 0 {
		println("Nil addr")
		return
	}
	net.Dial("tcp",c.addr)
}

func main() {
	client := NewClient()
	client.setAddr("localhost:8080")
	client.Connect()
}