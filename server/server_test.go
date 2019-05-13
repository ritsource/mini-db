package server_test

import (
	"testing"

	"github.com/ritwik310/mini-db/server"
	"github.com/ritwik310/mini-db/src"
)

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func HandleMsg(t *testing.T) {
	// Testing SET-Command hanlding
	bs1 := server.HandleMsg([]byte("SET\r\nkey1\r\n+OK\r\n"))
	d1, err := src.UnmarshalData(bs1)
	handleErr(t, err)

	if d1["error"] != nil {
		t.Error("d1[\"error\"] != nil")
	}

	// Testing GET-Command hanlding
	bs2 := server.HandleMsg([]byte("GET\r\nkey1\r\n"))
	d2, err := src.UnmarshalData(bs2)
	handleErr(t, err)

	if d2["error"] != nil {
		t.Error("d2[\"error\"] != nil")
	} else if d2["data"] != "OK" {
		t.Error("d2[\"data\"] != \"OK\"")
	}

	// Testing DELETE-Command
	bs3 := server.HandleMsg([]byte("DELETE\r\nkey1\r\n"))
	d3, err := src.UnmarshalData(bs3)
	handleErr(t, err)

	if d3["error"] != nil {
		t.Error("d3[\"error\"] != nil")
	}

	// Testing SET-Cmd for non existing key
	bs4 := server.HandleMsg([]byte("GET\r\nkey1\r\n"))
	d4, err := src.UnmarshalData(bs4)
	handleErr(t, err)

	if d4["error"] == nil {
		t.Error("d4[\"error\"] == nil", d4["error"], "or unable to delete file")
	}
}
