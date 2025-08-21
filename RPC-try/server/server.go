package main

import (
	"log"
	"net"
	"net/rpc"
	"strconv"
)

type Args struct {
	A,B int
}

type Adder struct {}

func (adder *Adder)Add(args * Args,result * int) error {
	*result = args.A + args.B
	log.Println(strconv.Itoa(args.A) + " + " + strconv.Itoa(args.B) + " = " + strconv.Itoa((*result)))
	return nil
}

func main() {
	adder := Adder{}
	rpc.Register(&adder)
	ln,_ := net.Listen("tcp","localhost:8080")
	for {
		conn,_ := ln.Accept()
		go rpc.ServeConn(conn)
	}
}