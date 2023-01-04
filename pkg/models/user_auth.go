package models

import "net"

type UserConn struct {
	net.Conn
	Auth     bool
	Id       int
	Username string
}

func (u *UserConn) UserAuth(username string) bool {
	u.Auth = true
	return true
}

func (u *UserConn) Authenticated() bool {
	return u.Auth
}
