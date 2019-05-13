package src

import (
	"bytes"
	"strconv"
)

// HandleProtocol parses message recieved from client
func HandleProtocol(bs []byte) (string, string, []interface{}, error) {
	var vals []interface{}
	var err error

	bss := bytes.Split(bs, []byte("\r\n"))

	cmd := bss[0]   // Command, GET / SET / DELETE
	key := bss[1]   // Key
	data := bss[2:] // Slice of Values

	// Iretating because, value only be parsed if it contains one
	for _, bs := range data {
		if len(bs) > 0 {
			switch bs[0] {
			case byte('+'):
				vals = append(vals, string(bs[1:])) // For "+" (string type)
			case byte(':'):
				val, err := strconv.Atoi(string(bs[1:])) // For ":" (interger type)
				if err != nil {
					break
				}
				vals = append(vals, val)
			case byte('$'):
				vals = append(vals, bs[1:]) // For "-" (binary / []byte type)
			}
		}
	}

	if err != nil {
		return "", "", vals, err
	}

	return string(cmd), string(key), vals, err
}
