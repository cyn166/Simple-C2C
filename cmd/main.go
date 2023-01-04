package main

import (
	"log"
	"net"

	"github.com/cyn166/Simple-C2C/pkg/models"
	"github.com/cyn166/Simple-C2C/pkg/telnet"
)

func main() {

	addr := net.JoinHostPort("127.0.0.1", "8023")
	log.Printf("Listening on %s", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, _ := listener.Accept()
		Userconn := &models.UserConn{Conn: conn, Auth: false}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		go telnet.Serve(Userconn)
	}
}
