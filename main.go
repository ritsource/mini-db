package main

import (
	"fmt"
	"strings"
)

func main() {
	// server.Start()
	// shell.Start()

	sl := strings.Split("GET\r\nkey\r\n", "\r\n")
	fmt.Println("S", sl[2:])
}
