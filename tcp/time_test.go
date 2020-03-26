package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMillisecond(t *testing.T) {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Millisecond)
	fmt.Println(int64(time.Millisecond))
	fmt.Println(time.Now().UnixNano() / int64(time.Millisecond))
}
