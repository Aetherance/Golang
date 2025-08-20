package main

var (
	client * Client
)

func main() {
	client = NewClient()
	client.setAddr("localhost:8080")
	client.Connect()

	menu()
}