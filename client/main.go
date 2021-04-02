package main

import (
	"fmt"
	"github.com/TStarFall/golangIMSystem/client/handler"
	"os"
)

var userId string
var userPwd string
var userName string

func main() {


	var key int
	for {
		fmt.Println("----------------欢迎登陆多人聊天室----------------")
		fmt.Println("1、登陆聊天系统")
		fmt.Println("2、注册用户")
		fmt.Println("3、退出系统")
		fmt.Print("请选择(1-3):")

		fmt.Scanln(&key)

		switch key {
		case 1:
			//提示用户输入用户id和用户pwd
			fmt.Println("请输入用户id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户pwd")
			fmt.Scanln(&userPwd)
			user := &handler.UserHandler{}
			user.Login(userId,userPwd)
		case 2:
			//提示用户输入用户id和用户pwd
			fmt.Println("请输入用户id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户pwd:")
			fmt.Scanln(&userPwd)
			fmt.Println("请输入用户nickName:")
			fmt.Scanln(&userName)
			user := &handler.UserHandler{}
			user.Register(userId,userPwd,userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("请输入(1-3)")
		}
	}

}
