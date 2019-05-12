package server_test

import (
	"reflect"
	"testing"

	"github.com/ritwik310/mini-db/server"
)

func TestProtocolHandler(t *testing.T) {
	// Testing for String
	k1, v1, err := server.HandleProtocol([]byte("#key1\r\n+OK\r\n"))
	if err != nil {
		t.Error(err)
	}
	if k1 != "key1" || v1 != "OK" {
		t.Errorf("HandleProtocol String: value mismatch %v != \"key1\" or %v != \"OK\"\n", k1, v1)
	}

	// Testing for Integer
	k2, v2, err := server.HandleProtocol([]byte("#key2\r\n:80\r\n"))
	if err != nil {
		t.Error(err)
	}
	if k2 != "key2" || v2 != 80 {
		t.Errorf("HandleProtocol Integer: value mismatch %v != \"key2\" or %v != \"80 (int)\"\n", k2, v2)
	}

	// Testing for Binary
	k3, v3, err := server.HandleProtocol([]byte("#key3\r\n$sliceofbytes\r\n"))
	if err != nil {
		t.Error(err)
	}
	if k3 != "key3" || !reflect.DeepEqual(v3, []byte("sliceofbytes")) {
		t.Errorf("HandleProtocol Binary: value mismatch %v != \"key3\" or %v-Error\n", k2, v2)
	}

}
