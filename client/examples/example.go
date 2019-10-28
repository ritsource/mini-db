package main

import (
	"fmt"

	"github.com/ritcrap/mini-db/client"
)

func main() {
	mdb := client.New("tcp", "localhost:8080")

	resp1, err := mdb.Set("myname", "Ritwik Saha", "str")
	resp2, err := mdb.Get("myname")
	resp3, err := mdb.Delete("myname")
	resp4, err := mdb.Get("myname")

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("resp1: %+v\n", resp1)
	fmt.Printf("resp2: %+v\n", resp2)
	fmt.Printf("resp3: %+v\n", resp3)
	fmt.Printf("resp4: %+v\n", resp4)
}
