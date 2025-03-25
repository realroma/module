package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]int{
		"Some": 1,
	}

	v, _ := json.Marshal(m)

	fmt.Println(v)

	_ = json.Unmarshal(v, &m)

	fmt.Println(m)
}
