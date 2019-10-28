package src_test

import (
	"reflect"
	"testing"

	"github.com/ritcrap/mini-db/src"
)

func TestInMemoryStore(t *testing.T) {
	store := src.Store{Persist: false}
	store.Map = make(map[string]interface{})

	// Tesing Set method
	store.Set("key1", "value1")
	if store.Map["key1"] != "value1" {
		t.Error("Set: value mismatch", store.Map["key1"], "!= value1")
	}

	// Testing Get method
	val, err := store.Get("key1")
	if err != nil {
		t.Error(err)
	}
	if val != "value1" {
		t.Error("Get: value mismatch", val, "!= value1")
	}

	// Testing Delete method
	err = store.Delete("key1")
	if err != nil {
		t.Error(err)
	}
	if store.Map["key1"] != nil {
		t.Error("Delete: value hasn't been deleted", store.Map["key1"], "!= nil")
	}

	// Testing MSet method
	wmap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	store.MSet(wmap)
	if store.Map["key1"] == nil || store.Map["key2"] == nil {
		t.Error("MSet: value is nil", "key1 =>", store.Map["key1"], "key2 =>", store.Map["key2"])
	}

	// Testing MGet method
	rmap := store.MGet([]string{"key1", "key2"})
	if !reflect.DeepEqual(wmap, rmap) {
		t.Errorf("MGet: value mismatch \n%+v\n%+v\n", wmap, rmap)
	}

	// Testing Flush method
	store.Set("key1", "value1")
	store.Set("key2", "value2")
	store.Flush()
	if store.Map["key1"] != nil || store.Map["key2"] != nil {
		t.Error("Flush: value hasn't been flushed", store.Map["key1"], "and", store.Map["key2"], "!= nil")
	}
}
