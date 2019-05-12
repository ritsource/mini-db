package main

import "github.com/ritwik310/mini-db/server"

func main() {
	server.Start()
	// fmt.Println(">")
	// server.Start()
	// cmd, key, val, err := src.HandleProtocol([]byte("SET\r\nkey1\r\n+OK\r\n"))
	// // cmd, key, val, err := src.HandleProtocol([]byte("GET\r\nkey1\r\n+OK\r\n"))
	// if err != nil {
	// 	panic(err)
	// }

	// var query interface{}
	// fmt.Println("Parts:", cmd, key, val)

	// store := src.Store{Persist: false}
	// store.Map = make(map[string]interface{})

	// switch cmd {
	// case "SET":
	// 	err = store.Set(key, val)
	// case "GET":
	// 	query, err = store.Get(key)
	// case "DELETE":
	// 	err = store.Delete(key)
	// }

	// fmt.Println(store.Map["key1"], query)

	// server.HandleProtocol([]byte("DELETE\r\nkey1\r\n+OK\r\n"))
	// server.HandleProtocol([]byte("HSET\r\nkey1\r\n+OK\r\n"))
	// server.HandleProtocol([]byte("HGET\r\nkey1\r\n+OK\r\n"))
}
