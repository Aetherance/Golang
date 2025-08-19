package main

import (
	"fmt"
	"os"
)

func menu() {
	fmt.Println("1. 登录")
	fmt.Println("2. 注册")
	fmt.Println("0. 退出")
	fmt.Print("> ")
	choose := ""
	fmt.Scan(choose)

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

func login() {
	username := ""
	password := ""
	fmt.Println("用户名: ")
	fmt.Scan(username)
	fmt.Scan(password)
}

func register() {

}