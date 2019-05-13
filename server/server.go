package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"

	"github.com/ritwik310/mini-db/src"
)

var store src.Store

func init() {
	// Initializing store
	store = src.Store{Persist: false}
	store.Map = make(map[string]interface{})
}

// Start starts a tcp server on specified port, default :8080.
// A client can send write to the TCP connection to manipulate data
func Start() error {
	// Starting TCP-Server
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Server Started...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConnection(conn) // Connection requests
	}
}

// handleConnection takes care of reading data from connection
// and writing the appropriate message (Error, Query)
func handleConnection(conn net.Conn) {
	// Reading connection data
	bs := make([]byte, 512)
	_, err := conn.Read(bs)
	if err != nil {
		conn.Write([]byte("Error: " + err.Error()))
		fmt.Println("Error:", err)
		return
	}

	// If message exists on connection
	if len(bs) > 0 {
		wbs := HandleMsg(bs) // Handling data from client, wbs => writable data

		// Writing response data (Error, Query)
		_, err = conn.Write(wbs)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	conn.Close() // Closing connection
}

// HandleMsg parses message data passed by client
// and does the specified query or insert or ...
func HandleMsg(bs []byte) []byte {
	// Parsing message data
	cmd, key, val, err := src.HandleProtocol(bs)
	if err != nil {
		return []byte("Error: " + err.Error())
	}

	// Doing the specified task in the message
	switch cmd {
	case "SET":
		// If SET-cmd
		err = store.Set(key, val)
		if err != nil {
			return []byte("Error: " + err.Error())
		}
	case "GET":
		// If GET-cmd
		newval, err := store.Get(key)
		if err != nil {
			return []byte("Error: " + err.Error())
		}

		// Converting data into []bytes to return
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		if err := enc.Encode(newval); err != nil {
			return []byte("Error: " + err.Error())
		}

		return buf.Bytes()

	case "DELETE":
		// If DELETE-cmd
		err = store.Delete(key)
		if err != nil {
			return []byte("Error: " + err.Error())
		}
	default:
		return []byte("Error: command not found")
	}

	return nil
}
