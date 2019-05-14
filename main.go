package main

import (
	"github.com/ritwik310/mini-db/server"
)

func main() {
	// cmd.Execute()
	// shell.Start()
	server.Start("8080", true, 5, "mydata.out")
}
