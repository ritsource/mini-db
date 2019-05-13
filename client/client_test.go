package client_test

import (
	"testing"

	"github.com/ritwik310/mini-db/client"
	"github.com/ritwik310/mini-db/server"
)

func Test(t *testing.T) {
	go server.Start()

	mdb := client.New("tcp", "localhost:8080")

	resp1, err := mdb.Set("myname", "Ritwik Saha", "str")
	resp2, err := mdb.Get("myname")
	resp3, err := mdb.Delete("myname")
	resp4, err := mdb.Get("myname")

	if err != nil {
		t.Error("Error:", err)
	}

	if resp1["status"] != 200 {
		t.Error("Error: status != 200")
	}

	if resp2["status"] != 200 || resp2["data"] != "Ritwik Saha" {
		t.Errorf("Error: status == %v, or data  == %v\n", resp2["status"], resp2["data"])
	}

	if resp3["status"] != 200 {
		t.Error("Error: status != 200")
	}

	if resp4["status"] != 400 || resp4["data"] != nil || resp4["error"] != "not found" {
		t.Errorf("Error: status == %v, or data  == %v, or error == %v\n", resp4["status"], resp4["data"], resp4["error"])
	}

}
