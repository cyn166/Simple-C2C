package login

import (
	"io"

	"github.com/cyn166/Simple-C2C/pkg/database"
	"github.com/cyn166/Simple-C2C/pkg/models"
	"github.com/ebarkie/telnet"
)

func Login(connUser *models.UserConn, tn *telnet.Ctx) {
	buf := make([]byte, 1024)
	tn.Write([]byte("Login\r\n"))
	tn.Write([]byte("Username: "))
	n, err := tn.Read(buf)
	if err == io.EOF {
		return
	}
	username := string(buf[:n])
	tn.Write([]byte("Password: "))
	n, err = tn.Read(buf)
	if err == io.EOF {
		return
	}
	password := string(buf[:n])

	user, err := database.GetUser(username, password)
	if err != nil {
		tn.Write([]byte("Invalid username or password\r\n"))
		return
	}
	connUser.Id = int(user.ID)
	connUser.Username = user.Username
	connUser.Auth = true
}
