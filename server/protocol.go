package server

import (
	"bytes"
	"strconv"
)

// var store

// func init() {
// 	store := src.Store{Persist: false}
// 	store.Map = make(map[string]interface{})
// }

// HandleProtocol parses event message
func HandleProtocol(bs []byte) (string, interface{}, error) {
	var val interface{}
	var err error
	bss := bytes.Split(bs, []byte("\r\n"))
	key := bss[0][1:]

	switch bss[1][0] {
	case byte('+'):
		val = handleStr(bss[1][1:])
	case byte(':'):
		val, err = handleInt(bss[1][1:])
	case byte('$'):
		val = handleBin(bss[1][1:])
	}

	if err != nil {
		return "", "", err
	}

	return string(key), val, err
}

func handleStr(bs []byte) string {
	return string(bs)
}

func handleInt(bs []byte) (int, error) {
	return strconv.Atoi(string(bs))
}

func handleBin(bs []byte) []byte {
	return bs
}
