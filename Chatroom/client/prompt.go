package main

import (
	"chatroom/Message"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func menu() {
	for {
		cmd := exec.Command("clear")
		cmd.Run()
		fmt.Println("1. 登录")
		fmt.Println("2. 注册")
		fmt.Println("0. 退出")
		fmt.Print("> ")
		choose := ""
		fmt.Scan(&choose)
	
		switch choose {
		case "1":
			login() 
		case "2":
			register()	
		case "0":
			os.Exit(0)
		default:
			fmt.Println("您只能输入 1,2,0 其中的一个!")
		}
	}
}

func login() {
	username := ""
	password := ""
	fmt.Println("用户名: ")
	fmt.Scan(&username)
	fmt.Println("密码: ")
	fmt.Scan(&password)

	if client.RequestLogin(username,password) {
		fmt.Println("登录成功!")
		chat()	
	} else {
		fmt.Println("登录失败!")
	}
}

func register() {
	username := ""
	password := ""
	fmt.Println("用户名: ")
	fmt.Scan(&username)
	fmt.Println("密码: ")
	fmt.Scan(&password)

	if client.RequestRegister(username,password) {
		fmt.Println("注册成功!")
		chat()
	} else {
		fmt.Println("注册失败!  已经有一个相同的用户名存在!")
	}
}

func chat() {
	clear()
	fmt.Println("Let' chat!")
	input := ""
	go client.recvLoop()
	for {
		fmt.Scan(&input)
		msg := message.SerMessage(client.LNam,input)
		m := map[string]string{}
		m["action"] = "Message"
		m["message"] = string(msg)
		str,_ := json.Marshal(m)
		client.send(string(str))
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}