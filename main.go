package main

import (
	"fmt"
	"reflect"

	"github.com/ritwik310/mini-db/src"
)

func main() {
	fmt.Println("Hello World!")
	wmap := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}

	err := src.WriteFile("data.out", wmap)
	if err != nil {
		panic(err)
	}

	rmap, err := src.ReadFile("data.out")
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(wmap, rmap) {
		panic("MGet: value mismatch")
	}
}
