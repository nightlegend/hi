package core

import (
	"encoding/json"
	"log"
	"net"

	"github.com/nightlegend/hi/dataconversion"
	"github.com/nightlegend/hi/errorpkg"
	"github.com/nightlegend/hi/handlers"
	"github.com/nightlegend/hi/lib/stringid"
	"github.com/nightlegend/hi/protocol"
)

// handler declaration.
var (
	userHandler    handlers.UsersHandler
	messageHandler handlers.MessageHandler
)

// SocketServer start a socket server.
func SocketServer() {
	var uid string
	conns := make(map[string]net.Conn)
	messages := make(chan string, 10)
	listener, err := net.Listen("tcp", ":9090")
	errorpkg.CheckErrors(err)
	defer listener.Close()
	go sendMsg(&conns, messages)
	for {
		conn, err := listener.Accept()
		errorpkg.CheckErrors(err)
		uid = stringid.GenerateNonCryptoID()
		conns[uid] = conn
		go handler(uid, conn, conns, messages)
	}
}

// handler is handle comming request.
func handler(uid string, conn net.Conn, conns map[string]net.Conn, message chan string) {
	log.Printf("new boy join: %v\n", conn.RemoteAddr().String())
	buff := make([]byte, 1024)
	for {
		length, err := conn.Read(buff)
		// errorpkg.CheckErrors(err)
		if err != nil {
			log.Println(err.Error())
			break
		}
		if length > 0 {
			buff[length] = 0
		}
		go dispatch(uid, conn, conns, buff[0:length])
		reviveMsg := string(buff[0:length])
		message <- reviveMsg
	}
}

// broadcast message
func sendMsg(conns *map[string]net.Conn, messages chan string) {
	for {
		rmsg := <-messages
		log.Println("Trans messages is: ", rmsg)
		// everyone will recived message.
		// for key, value := range *conns {
		// 	log.Println("new message from:", key)
		// 	_, err := value.Write([]byte(rmsg))
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 		value.Close()
		// 		delete(*conns, key)
		// 	}
		// }
	}
}

// dispatch request to different handler
func dispatch(uid string, conn net.Conn, conns map[string]net.Conn, recMsg []byte) {
	var tcpReq dataconversion.TCPRequest
	json.Unmarshal(recMsg, &tcpReq)
	switch tcpReq.Hd.HandlerID {
	case protocol.USER:
		// Go to user handler
		userHandler.MainHandler(uid, conn, tcpReq)
	case protocol.MESSAGE:
		// Go to chat message handler
		messageHandler.MainHandler(conn, conns, tcpReq)
	}
}
