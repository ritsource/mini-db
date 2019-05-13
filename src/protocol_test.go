package src_test

import (
	"reflect"
	"testing"

	server "github.com/ritwik310/mini-db/src"
)

func TestProtocolHandler(t *testing.T) {
	// Testing for String
	c1, k1, vs1, err := server.HandleProtocol([]byte("SET\r\nkey1\r\n+OK\r\n"))
	checkReturns(t, "HandleProtocol SET String", c1, "SET", k1, "key1", vs1, "OK", err)

	// Testing for Integer
	c2, k2, vs2, err := server.HandleProtocol([]byte("SET\r\nkey2\r\n:80\r\n"))
	checkReturns(t, "HandleProtocol SET Integer", c2, "SET", k2, "key2", vs2, 80, err)

	// Testing for Binary
	c3, k3, vs3, err := server.HandleProtocol([]byte("SET\r\nkey3\r\n$sliceofbytes\r\n"))
	checkReturns(t, "HandleProtocol SET Binary", c3, "SET", k3, "key3", vs3, []byte("sliceofbytes"), err)

	// Testing for String
	c4, k4, vs4, err := server.HandleProtocol([]byte("GET\r\nkey4\r\n"))
	checkReturns(t, "HandleProtocol SET String", c4, "GET", k4, "key4", vs4, nil, err)
}

func checkReturns(t *testing.T, msg string, cmd, cmdE string, key, keyE string, vals, valsE interface{}, err error) {
	if err != nil {
		t.Error(err)
	}

	if cmd != cmdE {
		t.Errorf("%v: command value mismatch %v != %v\n", msg, cmd, cmdE)
	}

	if key != keyE {
		t.Errorf("%v: key value mismatch %v != %v\n", msg, key, keyE)
	}

	if !reflect.DeepEqual(vals, valsE) {
		t.Errorf("%v: vals value mismatch \n%+v\n != %+v\n", msg, vals, valsE)
	}
}
