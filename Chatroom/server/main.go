package main

func main() {
	server := NewServer()
	server.Listen(":8080")
	server.Run()
}