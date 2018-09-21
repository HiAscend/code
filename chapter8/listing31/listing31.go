// This sample program demonstrates how to marshal a JSON string.
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Create a map of key/value pairs.
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// Marshal the map into a JSON string.
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))

	bytes, e := json.Marshal(c)
	if e != nil {
		log.Println("ERROR:", e)
		return
	}
	fmt.Println(string(bytes))

}
