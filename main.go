package main

import "github.com/ritwik310/mini-db/server"

func main() {
	// fmt.Println(">")
	// server.Start()
	server.HandleProtocol([]byte("#key1\r\n+OK\r\n"))
}
