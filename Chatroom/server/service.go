package main

import (
	"context"
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
	case "Message":
		s.onMessage(message)
	default:
		log.Println("parse error")
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

	val,_ := s.redis.HGet(context.Background(),"nameHashPswd",username).Result()
	if val == passwd {
		s.send("OK",conn)
		s.redis.SAdd(context.Background(),"onlineUser",username)
		s.userHashConn[username] = conn
	} else {
		s.send("FAIL",conn)
	}
}

func (s *Server)onRegister(username string,passwd string,conn net.Conn) {
	log.Println("A client is registering")

	exist,_ := s.redis.HExists(context.Background(),"nameHashPswd",username).Result()
	
	if exist {
		s.send("FAIL",conn)
	} else {
		s.redis.HSet(context.Background(),"nameHashPswd",username,passwd)	
		s.send("OK",conn)
	}
}

func (s *Server)onMessage(message string) {
	users,_ := s.redis.SMembers(context.Background(),"onlineUser").Result()
	for _,user := range users {
		conn := s.userHashConn[user]
		s.send(message,conn)
	}
}