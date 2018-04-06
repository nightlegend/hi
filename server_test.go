package main

import (
	"net"
	"testing"

	"github.com/nightlegend/hi/core"
)

func TestMain(t *testing.T) {
	go core.SocketServer()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:9090")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("socket server normally..")
	}
	conn.Close()
}
