package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	c := MyMessage{
		Path: "json",
		Body: "content",
	}
	bob, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(bob))
}
