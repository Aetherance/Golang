package main

import (
	"net/rpc"
	"os"
)

type Args struct {
	A,B int
}

func main() {
	client,err := rpc.Dial("tcp","localhost:8080")
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	var reply int
	client.Call("Adder.Add",&Args{2313,3123},&reply)
	
	println(reply)
}