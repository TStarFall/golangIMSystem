package main

import (
	"fmt"
	"github.com/TStarFall/golangIMSystem/server/model"
	"net"
	"time"
)

func process(conn net.Conn) {
	//延时关闭连接
	defer conn.Close()

	//
	handler := &HandlerConn{
		Conn: conn,
	}
	err := handler.ReadUserHandler()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误 err=", err)
		return
	}
}

//对UserDao进行初始化
func initUserDao() {
	//这里的pool是一个全局变量
	//
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//初始化顺序不能乱
	//初始化redis连接池
	initPool("localhost:6379",16,0,300 * time.Second)
	//初始化UserDao
	initUserDao()
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=",err)
		return
	}
	//一旦监听成功,就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept() err=", err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}

}
