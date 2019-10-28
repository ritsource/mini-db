package shell_test

import (
	"testing"

	"github.com/ritcrap/mini-db/shell"
)

func TestFormatCmd(t *testing.T) {
	assertionMap := map[string]string{
		"SET key1 value1":       "SET\r\nkey1\r\n+value1\r\n",
		"SET key1 Ritwik Saha":  "SET\r\nkey1\r\n+Ritwik\r\n+Saha\r\n",
		"SET key1 value1 --str": "SET\r\nkey1\r\n+value1\r\n",
		"SET key1 80 --integer": "SET\r\nkey1\r\n:80\r\n",
		"SET key1 88 --int":     "SET\r\nkey1\r\n:88\r\n",
		"SET key1 value1 --bin": "SET\r\nkey1\r\n$value1\r\n",
		"GET key1":              "GET\r\nkey1\r\n",
		"DELETE key1":           "DELETE\r\nkey1\r\n",
	}

	for k, v := range assertionMap {
		fstr := shell.FormatCmd(k)
		if fstr != v {
			t.Errorf("Error: in %v\n %v != %v\n", k, fstr, v)
		}
	}
}
