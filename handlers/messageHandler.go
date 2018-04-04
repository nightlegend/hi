package handlers

import (
	"encoding/json"
	"net"

	"github.com/nightlegend/hi/dataconversion"
)

// ChatMessage is define trans message struct
type ChatMessage struct {
	ID          string
	ChatType    string
	FromID      string
	TOID        string
	GroupName   string
	MessageType string
	Message     []byte
	UserList    []UserInfo
}

// transUserMessage transform user message(one to one)
func transUserMessage(conn net.Conn, conns map[string]net.Conn, req dataconversion.TCPRequest) {
	var chatMsg ChatMessage
	var resp dataconversion.TCPResponse
	// var header dataconversion.Header
	var session net.Conn
	json.Unmarshal(req.Body, &chatMsg)
	session = conns[chatMsg.TOID]
	resp = dataconversion.TCPResponse{
		HD:   req.Hd,
		Body: req.Body,
	}
	data, _ := json.Marshal(resp)
	session.Write(data)
}
