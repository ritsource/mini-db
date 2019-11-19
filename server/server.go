package server

import (
	"errors"
	"fmt"
	"net"

	"github.com/ritsource/mini-db/src"
)

// Store instance that conatins all the in-memory data
var store src.Store

// Start starts a tcp server on specified port, default :8080.
// A client can send write to the TCP connection to manipulate data
func Start(port string, persist bool, delay int, output string) error {
	//  Initializing store
	store.Init(persist, delay, output)

	// Starting TCP-Server
	ln, err := net.Listen("tcp", ":"+port)
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
		wbs := HandleMsg(&store, bs) // Handling data from client, wbs => writable data

		// Writing response data (Error, Query)
		_, err = conn.Write(wbs)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	conn.Close() // Closing connection
}

// buildResponse builds response map to be sent
func buildResponse(st int, d interface{}, er error) []byte {
	resmap := map[string]interface{}{}

	resmap["status"] = st
	resmap["data"] = d
	if er != nil {
		resmap["error"] = er.Error()
	} else {
		resmap["error"] = nil
	}

	// Convert to []byte
	bs, err := src.MarshalData(&resmap)
	if err != nil {
		return []byte("Error: " + err.Error())
	}

	return bs
}

// HandleMsg parses message data passed by client
// and does the specified query or insert or ...
func HandleMsg(st *src.Store, bs []byte) []byte {
	// Parsing message data
	cmd, key, val, err := src.HandleProtocol(bs)
	if err != nil {
		return []byte("Error: " + err.Error())
	}

	// Doing the specified task in the message
	switch cmd {
	case "SET":
		// If SET-cmd
		err = st.Set(key, val)
		if err != nil {
			return buildResponse(400, nil, err)
		}
	case "GET":
		// If GET-cmd
		newval, err := st.Get(key)
		if err != nil {
			return buildResponse(400, nil, err)
		}
		return buildResponse(200, newval, nil)
	case "DELETE":
		// If DELETE-cmd
		err = st.Delete(key)
		if err != nil {
			return buildResponse(400, nil, err)
		}
	default:
		return buildResponse(400, nil, errors.New("Error: command not found"))
	}

	return buildResponse(200, nil, nil)
}
