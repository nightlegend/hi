package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"testing"

	"github.com/nightlegend/hi/core"
)

const (
	// LoginRequest login request flag
	LoginRequest = 0x0001
	// LoginRequest login success flag
	LoginSuccess = 0x0010
	// LoginFailed login failed flag
	LoginFailed = 0x0100

	UserMessageRequest = 0x0002
	UserMessageSuccess = 0x0020
	UserMessageFail    = 0x00200
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

// UserInfo user information type

type Header struct {
	HandlerID int
	CommandID int
}

type TCPRequest struct {
	Hd   Header
	Body []byte
}

// TCPResponse a tcp response struct
type TCPResponse struct {
	HD   Header
	Body []byte
}

type UserInfo struct {
	ID       string
	Name     string
	Password string
	Message  string
}

var userInfo UserInfo

func TestMain(t *testing.T) {
	go core.SocketServer()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "localhost:9090")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("socket server normally..")
	}
	go sendMsg(conn)
	buff := make([]byte, 1024)
	for {
		length, err := conn.Read(buff)
		if err != nil {
			conn.Close()
			t.Log("Server crashed..bye")
			os.Exit(0)
		}
		t.Log(string(buff[0:length]))
		go handler(conn, buff[0:length], t)
	}
}

// sendMsg send message to server.
func sendMsg(conn net.Conn) {
	var user UserInfo
	var tcpReq TCPRequest
	var header Header
	var data []byte
	user = UserInfo{
		ID:       "",
		Name:     "David Guo",
		Password: "12345",
		Message:  "Hi ,你好",
	}
	body, _ := json.Marshal(user)
	header = Header{
		HandlerID: 0x0001,
		CommandID: LoginRequest,
	}
	tcpReq = TCPRequest{
		Hd:   header,
		Body: body,
	}
	data, _ = json.Marshal(tcpReq)

	length, err := conn.Write(data)
	log.Println(length)
	if err != nil {
		log.Println(err.Error())
		conn.Close()
		os.Exit(0)
	}
}

// handler handle response.
func handler(conn net.Conn, data []byte, t *testing.T) {
	var resp TCPResponse
	var chatMessage ChatMessage
	json.Unmarshal(data, &resp)
	switch resp.HD.HandlerID {
	case 0x0001:
		json.Unmarshal(resp.Body, &userInfo)
		t.Log(userInfo.ID)
		os.Exit(0)
	case 0x0002:
		json.Unmarshal(resp.Body, &chatMessage)
		t.Log(string(chatMessage.Message))
		os.Exit(0)
	}
}
