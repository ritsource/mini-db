package src_test

import (
	"testing"

	"github.com/ritwik310/mini-db/src"
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
		t.Error("Set: value hasn't been deleted", store.Map["key1"], "is not nil")
	}

}
