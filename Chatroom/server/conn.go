package main

import (
	"net"
	"sync"
)

type Conns struct {
	data map[string]net.Conn
	lock sync.Mutex
}

func (conns *Conns)addConn(conn net.Conn) {
	conns.lock.Lock()
	conns.data[conn.RemoteAddr().String()] = conn
	conns.lock.Unlock()
}

func (conns *Conns)removeConn(conn net.Conn) {
	conns.lock.Lock()
	defer conns.lock.Unlock()
	delete(conns.data,conn.RemoteAddr().String())
}

func NewConns() * Conns {
	return &Conns{
		data: make(map[string]net.Conn),
	}
}