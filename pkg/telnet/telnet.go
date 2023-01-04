package telnet

import (
	"io"
	"log"

	"github.com/cyn166/Simple-C2C/pkg/commands"
	"github.com/cyn166/Simple-C2C/pkg/models"
	"github.com/ebarkie/telnet"
)

func Serve(conn *models.UserConn) {
	defer conn.Close()
	defer log.Printf("Connection from %s closed", conn.RemoteAddr())

	tn := telnet.NewReadWriter(conn)

	tn.Write([]byte("Welcome to a test telnet server!\r\n\r\n"))

	buf := make([]byte, 1024)
	// if conn.Auth {
	// 	tn.Write([]byte("auth:\r\n"))
	// }
	for {
		tn.Write([]byte("> "))
		n, err := tn.Read(buf)
		if err == io.EOF {
			return
		}
		cmd := string(buf[:n])

		commands.Setup(conn, cmd, tn)
	}
}
