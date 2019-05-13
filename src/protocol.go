package src

import (
	"bytes"
	"strconv"
)

// HandleProtocol parses message recieved from client
func HandleProtocol(bs []byte) (string, string, interface{}, error) {
	var val interface{}
	var err error

	bss := bytes.Split(bs, []byte("\r\n"))

	cmd := bss[0]   // Command, GET / SET / DELETE
	key := bss[1]   // Key
	rest := bss[2:] // Slice of Values

	if len(rest[0]) > 0 {
		switch rest[0][0] {
		case byte('+'):
			val = handleStr(&rest) // For "+" (string type)
		case byte(':'):
			val, err = strconv.Atoi(string(rest[0][1:])) // For ":" (interger type)
			break
		case byte('$'):
			val = handleBin(&rest) // For "$" (binary / []byte type)
		}
	}

	if err != nil {
		return "", "", val, err
	}

	return string(cmd), string(key), val, err
}

// handleStr
func handleStr(rest *[][]byte) string {
	var val string
	for i, r := range *rest {
		if i != len(*rest)-1 {
			if i == 0 {
				val += string(r[1:])
			} else {
				val += (" " + string(r[1:]))
			}
		}
	}

	return val
}

func handleBin(rest *[][]byte) []byte {
	var val []byte
	for i, r := range *rest {
		if i != len(*rest)-1 {
			if i == 0 {
				val = append(val, r[1:]...)
			} else {
				val = append(val, byte(' '))
				val = append(val, r[1:]...)
			}
		}
	}

	return val
}

// func handleInt() {

// }
