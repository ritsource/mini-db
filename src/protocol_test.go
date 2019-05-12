package src_test

import (
	"reflect"
	"testing"

	server "github.com/ritwik310/mini-db/src"
)

func TestProtocolHandler(t *testing.T) {
	// Testing for String
	c1, k1, v1, err := server.HandleProtocol([]byte("SET\r\nkey1\r\n+OK\r\n"))
	if err != nil {
		t.Error(err)
	}
	if c1 != "SET" && k1 != "key1" || v1 != "OK" {
		t.Errorf("HandleProtocol String: value mismatch %v != \"SET\" or %v != \"key1\" or %v != \"OK\"\n", c1, k1, v1)
	}

	// Testing for Integer
	c2, k2, v2, err := server.HandleProtocol([]byte("SET\r\nkey2\r\n:80\r\n"))
	if err != nil {
		t.Error(err)
	}
	if c2 != "SET" && k2 != "key2" || v2 != 80 {
		t.Errorf("HandleProtocol Integer: value mismatch %v != \"SET\" or %v != \"key2\" or %v != \"80 (int)\"\n", c2, k2, v2)
	}

	// Testing for Binary
	c3, k3, v3, err := server.HandleProtocol([]byte("SET\r\nkey3\r\n$sliceofbytes\r\n"))
	if err != nil {
		t.Error(err)
	}
	if c3 != "SET" && k3 != "key3" || !reflect.DeepEqual(v3, []byte("sliceofbytes")) {
		t.Errorf("HandleProtocol Binary: value mismatch %v != \"SET\" or %v != \"key3\" or %v-Error\n", c3, k2, v2)
	}

}
