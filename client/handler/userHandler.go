package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/client/utils"
	common "github.com/TStarFall/golangIMSystem/common"
	"net"
	"os"
)

type UserHandler struct {
	//暂时不需要字段
}

func MessageMarshal(tf *utils.Transfer,message *common.Message, v interface{})(mes common.Message, err error) {
	//序列化
	data, err := json.Marshal(v)
	//fmt.Println(v)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5、把data赋给mes.Data字段
	message.Data = string(data)

	//6、将mes序列化
	data, err = json.Marshal(message)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//7发送data给服务器

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("Register WritePkg(mes) err=",err)
		return
	}

	//8、回读服务器返回注册消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("Register ReadPkg err=",err)
		return
	}
	return mes, err
}

//定义登陆协议
func (this *UserHandler)Login(userId, userPwd string) (err error) {
	//验证userId和userPwd是否一致
	//fmt.Printf("userId=%s, userPwd=%s\n", userId, userPwd)
	//return nil

	//1、连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("net.Dial err=",err)
		return err
	}

	//2、准备通过conn发送消息给服务器
	var mes common.Message
	mes.Type = common.LoginMesType

	//3、创建一个LoginMes 结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//创建一个发送连接实例
	transfer := &utils.Transfer{
		Conn: conn,
	}
	//序列化，发送到服务器，并接收回传
	mes, err = MessageMarshal(transfer,&mes, loginMes)

	var loginResMes common.LoginResMes

	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	//fmt.Println("loginResMes=",loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = common.UserOnline

		//显示在线用户列表
		fmt.Println("当前在线用户列表：")
		//遍历loginResMes.UserId
		for _, v := range loginResMes.UserId {
			if v == userId {
				continue
			}
			fmt.Printf("用户id：%s\n", v)

			user := &common.User{
				UserId: v,
				UserStatus: common.UserOnline,
			}
			onlineUsers[v] = user
		}

		//该协程保持和服务器的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端终端
		go MessageHandler(transfer.Conn)
		ShowMenu()
	} else if loginResMes.Code ==500 {
		fmt.Println(loginResMes.Error)
	}
	return
}


func (this *UserHandler) Register(userId,userPwd,userName string)(err error) {
	conn, err := net.Dial("tcp","localhost:8889")
	defer conn.Close()

	if err != nil {
		fmt.Println("Register net.Dial err=",err)
		return
	}

	//2、准备通过conn发送消息给服务器
	var mes common.Message
	mes.Type = common.RegisterMesType
	//3、创建一个RegisterMes实例
	var register = common.RegisterMes{
		User: common.User{
			UserId: userId,
			UserPwd: userPwd,
			UserName: userName,
		},
	}
	//创建一个发送连接实例
	transfer := &utils.Transfer{
		Conn: conn,
	}
	//4.序列化register
	mes, err = MessageMarshal(transfer,&mes,register)

	//9、反序列化mes
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登陆")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
