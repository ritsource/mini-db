package src

import (
	"bytes"
	"errors"
	"strconv"
)

// Only in dev!
func checkFormat(bss [][]byte) error {
	if (string(bss[0]) == "SET" || string(bss[0]) == "HSET") && len(bss) < 3 {
		return errors.New("Not enough arguements in message")
	} else if len(bss) < 2 {
		return errors.New("Not enough arguements in message")
	}

	return nil
}

// HandleProtocol parses event message
func HandleProtocol(bs []byte) (string, string, interface{}, error) {
	var val interface{}
	bss := bytes.Split(bs, []byte("\r\n"))
	err := checkFormat(bss)
	if err != nil {
		return "", "", "", err
	}

	cmd := bss[0]
	key := bss[1]

	switch bss[2][0] {
	case byte('+'):
		val = string(bss[2][1:])
	case byte(':'):
		val, err = strconv.Atoi(string(bss[2][1:]))
	case byte('$'):
		val = bss[2][1:]
	}

	if err != nil {
		return "", "", "", err
	}

	return string(cmd), string(key), val, err
}
