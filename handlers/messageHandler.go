package handlers

import (
	"encoding/json"
	"log"
	"net"

	"github.com/nightlegend/hi/dataconversion"
)

// ChatMessage is define trans message struct
type ChatMessage struct {
	ID          string
	ChatType    string //1. Group 2.Signal
	FromID      string
	TOID        string
	GroupName   string
	MessageType string //1. wrods, 2.picture, 3.voice/vedio
	Message     []byte
	UserList    []UserInfo //if chat type is group, then here is store group user list.
}

// transUserMessage transform user message(one to one)
func transUserMessage(conn net.Conn, conns map[string]net.Conn, req dataconversion.TCPRequest) {
	var chatMsg ChatMessage
	var resp dataconversion.TCPResponse
	// var header dataconversion.Header
	var session net.Conn
	json.Unmarshal(req.Body, &chatMsg)
	log.Println(chatMsg.TOID)
	log.Println(chatMsg.Message)
	session = conns[chatMsg.TOID]
	resp = dataconversion.TCPResponse{
		HD:   req.Hd,
		Body: req.Body,
	}
	data, _ := json.Marshal(resp)
	session.Write(data)
}

// ack send a ack to client.
// TO-DO
// 1. define ACK struct
func ack(conn net.Conn, req dataconversion.TCPRequest) {
	var resp dataconversion.TCPResponse
	// TO-DO: logic of send ack.
	resp = dataconversion.TCPResponse{
		HD:   req.Hd,
		Body: nil,
	}
	data, _ := json.Marshal(resp)
	conn.Write(data)
}
