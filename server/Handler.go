package main

import (
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
	"github.com/TStarFall/golangIMSystem/server/handler"
	"github.com/TStarFall/golangIMSystem/server/utils"
	"io"
	"net"
)

type HandlerConn struct {
	Conn net.Conn
}

//根据客户发送的消息种类不同，决定调用哪个函数处理
func (this *HandlerConn)MessageHandler(mes *common.Message) (err error) {

	switch mes.Type {
	case common.LoginMesType:
		//处理登陆逻辑
		//创建一个用户连接实例
		userConn := handler.UserConn{
			Conn: this.Conn,
		}
		err = userConn.LoginHandler(mes)
		if err != nil {
			fmt.Println("userConn.LoginHandler(mes) err=", err)
			return
		}
	case common.RegisterMesType:
		//处理注册逻辑
		//创建一个用户连接实例
		userConn := handler.UserConn{
			Conn: this.Conn,
		}
		err = userConn.RegisterHandler(mes)
		if err != nil {
			fmt.Println("userConn.RegisterHandler(mes) err=", err)
			return
		}
	case common.SmsMesType:
		//创建一个SmsHandler实例转发消息
		smsHandler := &handler.SmsHandler{}
		smsHandler.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return err
}

//读取客户端发送的消息
func (this *HandlerConn) ReadUserHandler() (err error) {
	//循环读取客户端发送的信息
	for {
		transfer := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("等待下次输入...")
				return err
			}
			return err
		}
		err = this.MessageHandler(&mes)
		if err != nil {
			return err
		}
	}
}
