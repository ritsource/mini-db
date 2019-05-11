package main

import "fmt"

func main() {
	fmt.Println("Hwllo World!")

	x := map[string]interface{}{
		"n": "C",
	}

	fmt.Println("x[\"name\"]", x["name"])
}
