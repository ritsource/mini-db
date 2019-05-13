package src

import (
	"bytes"
	"strconv"
)

// HandleProtocol parses event message
func HandleProtocol(bs []byte) (string, string, []interface{}, error) {
	var vals []interface{}
	var err error

	bss := bytes.Split(bs, []byte("\r\n"))

	cmd := bss[0]
	key := bss[1]
	data := bss[2:]

	for _, bs := range data {
		if len(bs) > 0 {
			switch bs[0] {
			case byte('+'):
				vals = append(vals, string(bs[1:]))
			case byte(':'):
				val, err := strconv.Atoi(string(bs[1:]))
				if err != nil {
					break
				}
				vals = append(vals, val)
			case byte('$'):
				vals = append(vals, bs[1:])
			}
		}
	}

	if err != nil {
		return "", "", vals, err
	}

	return string(cmd), string(key), vals, err
}
