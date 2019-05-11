package src

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

// WriteFile writes a file and dumps given data into it
func WriteFile(fpath string, d map[string]interface{}) error {
	// Create file
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get Binary
	b, err := MarshalData(&d)
	if err != nil {
		return err
	}

	// Write to File
	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile reads a file that stores data and returns map[string]interface{}
func ReadFile(fpath string) (map[string]interface{}, error) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	var d map[string]interface{}
	d, err = UnmarshalData(b)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// MarshalData encodes map[string]interface into binary
func MarshalData(d *map[string]interface{}) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(*d); err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

// UnmarshalData reads []byte and decodes it into map[string]interface{}
func UnmarshalData(b []byte) (map[string]interface{}, error) {
	var data map[string]interface{}

	reader := bytes.NewReader(b)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
