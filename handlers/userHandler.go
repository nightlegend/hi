package handlers

import (
	"encoding/json"
	"log"
	"net"

	"github.com/nightlegend/hi/dataconversion"
	"github.com/nightlegend/hi/protocol"
)

// UserInfo is a user contain information.
type UserInfo struct {
	ID       string
	Name     string
	Password string
	Message  string
}

// Login handle login request
func login(uid string, conn net.Conn, request dataconversion.TCPRequest) {
	var userInfo UserInfo
	var resp dataconversion.TCPResponse
	var header dataconversion.Header
	json.Unmarshal(request.Body, &userInfo)
	log.Println(userInfo.Message)
	log.Println(userInfo.Name)
	// TO-DO
	// verification user info from db

	userInfo.ID = uid
	header = dataconversion.Header{
		HandlerID: protocol.USER,
		CommandID: protocol.LoginSuccess,
	}
	body, _ := json.Marshal(userInfo)
	resp = dataconversion.TCPResponse{
		HD:   header,
		Body: body,
	}
	data, _ := json.Marshal(resp)
	conn.Write(data)
}
