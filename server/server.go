package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"

	"github.com/ritwik310/mini-db/src"
)

// Start ...
func Start() error {
	// Starting TCP Server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Server Started...")

	// Looking for Events
	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		bs, err := ioutil.ReadAll(conn)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if len(bs) > 0 {
			handleEvent(conn, bs)
		}

		conn.Close()
	}
}

// handleEvent

func writeToClient(conn io.Writer, str string) {
	_, err2 := io.WriteString(conn, str)
	if err2 != nil {
		fmt.Println("Server Error:", err2)
	}
}

func handleEvent(conn io.Writer, bs []byte) {
	cmd, key, val, err := src.HandleProtocol(bs)

	// Writing error to the client
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var query interface{}
	fmt.Println("Parts:", cmd, key, val)

	store := src.Store{Persist: false}
	store.Map = make(map[string]interface{})

	switch cmd {
	case "SET":
		err = store.Set(key, val)
	case "GET":
		query, err = store.Get(key)
	case "DELETE":
		err = store.Delete(key)
	}

	fmt.Println("Result:", store.Map["key1"], query)
}
