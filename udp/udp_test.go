package main

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr(t *testing.T) {
	fmt.Print(net.Addr())
}