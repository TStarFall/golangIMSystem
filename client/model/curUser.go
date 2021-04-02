package model

import (
	"github.com/TStarFall/golangIMSystem/common"
	"net"
)

type CurUser struct {
	Conn net.Conn
	common.User
}
