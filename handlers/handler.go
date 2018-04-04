package handlers

import (
	"net"

	"github.com/nightlegend/hi/dataconversion"
	"github.com/nightlegend/hi/protocol"
)

// Handlers provide an interface for handler.
type Handlers interface {
	MainHandler()
}

// UsersHandler provide user (login, logout, etc..)handler functionality.
type UsersHandler struct {
	userInfo UserInfo
}

// MessageHandler provide chat message handler functionality.
type MessageHandler struct {
	chatMsg ChatMessage
}

// MainHandler handle user request.
func (userHandler *UsersHandler) MainHandler(uid string, conn net.Conn, req dataconversion.TCPRequest) {
	switch req.Hd.CommandID {
	case protocol.LoginRequest:
		login(uid, conn, req)
	}
}

// MainHandler handle chat message request.
func (msgHandler *MessageHandler) MainHandler(conn net.Conn, conns map[string]net.Conn, req dataconversion.TCPRequest) {
	switch req.Hd.CommandID {
	case protocol.UserMessageRequest:
		transUserMessage(conn, conns, req)
	}
}
