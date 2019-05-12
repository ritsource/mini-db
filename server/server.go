package server

import (
	"fmt"
	"io/ioutil"
	"net"
)

// Start ...
func Start() error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		// _, err = io.WriteString(conn, fmt.Sprintln("Hello World!"))
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// }

		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			panic(err)
		}
		if len(bs) > 0 {
			fmt.Printf("Result: %+s\n", bs)
		}

		conn.Close()
	}
}
