package src_test

import (
	"reflect"
	"testing"

	"github.com/ritsource/mini-db/src"
)

func TestDataMarshalAndUnmarshal(t *testing.T) {
	wmap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	// Testing MarshalData
	b, err := src.MarshalData(&wmap)
	if err != nil {
		t.Error(err)
	}

	// Testing UnmarshalData
	rmap, err := src.UnmarshalData(b)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(wmap, rmap) {
		t.Errorf("Marshal & Unmarshal: value mismatch \n%+v\n%+v\n", wmap, rmap)
	}
}

func TestWriteaAndReadFile(t *testing.T) {
	wmap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	err := src.WriteFile("data.out", wmap)
	if err != nil {
		t.Error(err)
	}

	rmap, err := src.ReadFile("data.out")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(wmap, rmap) {
		t.Errorf("ReadFile & WriteFile: value mismatch \n%+v\n%+v\n", wmap, rmap)
	}
}
