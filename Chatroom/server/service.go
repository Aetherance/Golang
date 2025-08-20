package main

import (
	// "github.com/redis/go-redis/v9"
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
)

func (s *Server)parseMessage(message string,conn net.Conn) {
	value := map[string]string{}
	json.Unmarshal([]byte(message),&value)
	action := value["action"]
	switch action {
	case "Login":
		s.onLogin(value["username"],value["passwd"],conn)
	case "Register":
		s.onRegister(value["username"],value["passwd"],conn)
	default:
		log.Println("解析错误")
	}
}

func (s *Server)send(msg string,conn net.Conn) {
	length := len(msg)
	buff := make([]byte,4)
	binary.BigEndian.PutUint32(buff,uint32(length))
	buff = append(buff, []byte(msg)...)
	net.Conn.Write(conn,buff)
}

func (s *Server)onLogin(username string,passwd string,conn net.Conn) {
	log.Println("A client is logging in")
	log.Println("Username: " + username + " Password: " + passwd)
	// database
	s.send("OK",conn)
}

func (s *Server)onRegister(username string,passwd string,conn net.Conn) {
	log.Println("A client is registering")
}