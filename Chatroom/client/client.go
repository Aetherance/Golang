package main

import (
	// "chatroom/Message"
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"net"
)

type Client struct {
	addr string
	conn net.Conn
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
	conn,_ := net.Dial("tcp",c.addr)
	c.conn = conn
}

func (c*Client) send(message string) {
	length := make([]byte,4)
	binary.BigEndian.PutUint32(length,uint32(len(message)))
	length = append(length, []byte(message)...)
	log.Println(string(length))
	c.conn.Write(length)
}

func (c*Client) recv() string {
	length := make([]byte,4)
	io.ReadFull(c.conn,length)
	Len := binary.BigEndian.Uint32(length)
	buff := make([]byte,Len)
	io.ReadFull(c.conn,buff)
	return string(buff)
}

func (c * Client) RequestLogin(username string,passwd string) bool {
	msg := map[string]string{}
	msg["action"] = "Login"
	msg["username"] = username
	msg["passwd"] = passwd
	str,_ := json.Marshal(msg)
	c.send(string(str))
	ret := c.recv()
	if ret == "OK" {
		return true
	} else {
		return false
	}
}

func (c * Client) RequestRegister() {

}