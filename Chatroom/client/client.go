package main

import (
	"chatroom/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Client struct {
	addr string
	conn net.Conn
	LNam string
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
		c.LNam = username
		return true
	} else {
		return false
	}
}

func (c * Client) RequestRegister(username string,passwd string) bool {
	msg := map[string]string{}
	msg["action"] = "Register"
	msg["username"] = username
	msg["passwd"] = passwd
	str,_ := json.Marshal(msg)
	c.send(string(str))
	ret := c.recv()
	if ret == "OK" {
		c.LNam = username
		return true
	} else {
		return false
	}
}

func (c * Client) recvLoop() {
	for {
		str := c.recv()
		val := map[string]string{}
		buff := []byte(str)
		json.Unmarshal(buff,&val)
		msg := val["message"]
		value := message.UnSerMessage([]byte(msg))
		fmt.Println(value.From,": ",value.Text)
	}
}