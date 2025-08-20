package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type Server struct {
	ln net.Listener
	conns Conns
	isListening bool
}

func NewServer() * Server {
	return &Server{
		conns: *NewConns(),
		isListening: false,
	}
}

func (s * Server)Run() {
	go s.GetCommand()
	for {
		conn,_ := s.ln.Accept()
		log.Println("New conn from: " + conn.RemoteAddr().String())
		s.conns.addConn(conn)
		go s.handleConn(conn)
	}
}

func (s * Server)Stop() {
	s.ln.Close()
}

func (s *Server)Listen(addr string) {
	s.ln,_ = net.Listen("tcp",addr)
	log.Println("Server listening " + addr)
	s.isListening = true
}

func (s * Server)GetCommand() {
	for {
		cmd := ""
		fmt.Scan(&cmd)
		if cmd == "stop" {
			if s.isListening {
				s.Stop()
			} else {
				fmt.Println("Server: Command Not Fount!")
			}
			log.Print("Exiting gracefully...")
			os.Exit(0)
		}
	}
}

func (s *Server)handleConn(conn net.Conn) {
	for {
		lenBuf := [4]byte{}
		_,err := io.ReadFull(conn,lenBuf[:])
		
		if err == io.EOF {
			log.Println("Client " + conn.RemoteAddr().String() + " left!")
			return
		}
		
		length := binary.BigEndian.Uint32(lenBuf[:])
		buff := make([]byte,length)
		_,err = io.ReadFull(conn,buff)

		if err == io.EOF {
			log.Println("Client " + conn.RemoteAddr().String() + " left!")
			return
		}

		log.Println(string(buff))

		s.parseMessage(string(buff),conn)
	}
}