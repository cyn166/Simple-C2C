package commands

import (
	"github.com/cyn166/Simple-C2C/pkg/commands/addclient"
	"github.com/cyn166/Simple-C2C/pkg/commands/login"
	"github.com/cyn166/Simple-C2C/pkg/database"
	"github.com/cyn166/Simple-C2C/pkg/models"
	"github.com/ebarkie/telnet"
)

func Setup(connUser *models.UserConn, cmd string, tn *telnet.Ctx) {
	switch cmd {
	case "quit\r\n":
		connUser.Close()
	case "help\r\n":

		if !checkAdmin(connUser, tn) {
			return
		}
		tn.Write([]byte("Commands:\r\n"))
		tn.Write([]byte("  help - display this help\r\n"))
		tn.Write([]byte("  login - login\r\n"))
		tn.Write([]byte("  quit - close the connection\r\n"))
	case "login\r\n":
		login.Login(connUser, tn)
	default:
		tn.Write([]byte("Unknown command\r\n"))
	}

}

func checkAuth(connUser *models.UserConn, tn *telnet.Ctx) {
	if !connUser.Auth {
		tn.Write([]byte("Please login first\r\n"))
		err := connUser.Close()
		if err != nil {
			tn.Write([]byte("Error closing listener\r\n"))
		}
		return
	}
}

func checkAdmin(connUser *models.UserConn, tn *telnet.Ctx) bool {
	var user models.User
	db := database.GetDB()
	db.First(&user, "id = ?", connUser.Id)
	if !user.Admin {
		tn.Write([]byte("You are not an admin\r\n"))
		return false
	}
	return true
}
