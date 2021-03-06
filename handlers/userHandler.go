package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/nightlegend/hi/dataconversion"
	"github.com/nightlegend/hi/protocol"
	redisUtils "github.com/nightlegend/hi/utils"
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
	log.Println("User Name:", userInfo.Name)
	log.Println("User Comment:", userInfo.Message)
	// TO-DO
	// 1. verification user info from db.
	// 2. save session to redis.
	cli, err := redisUtils.NewCli()
	if err != nil {
		// To-do item
		// save user in someplace
	} else {
		connStr := fmt.Sprintf("%v", conn)
		redisUtils.Set(cli, uid, connStr)
		redisUtils.Close(cli)
	}
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

// logout handle logout request
func logout(conn net.Conn, request dataconversion.TCPRequest) {
	var userInfo UserInfo
	json.Unmarshal(request.Body, &userInfo)
	cli, err := redisUtils.NewCli()
	if err != nil {
		log.Fatal(err)
	} else {
		redisUtils.Delete(cli, userInfo.ID)
	}
	header := dataconversion.Header{
		HandlerID: protocol.USER,
		CommandID: protocol.LogoutSuccess,
	}
	body, _ := json.Marshal(userInfo)
	resp := dataconversion.TCPResponse{
		HD:   header,
		Body: body,
	}
	data, _ := json.Marshal(resp)
	conn.Write(data)
}
