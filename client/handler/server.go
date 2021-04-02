package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/client/utils"
	"github.com/TStarFall/golangIMSystem/common"
	"net"
	"os"
)

func ShowMenu() {
	for {
		fmt.Println("----------恭喜xxx登陆成功----------")
		fmt.Println("----------1、显示在线用户列表----------")
		fmt.Println("----------2、发送消息----------")
		fmt.Println("----------3、信息列表----------")
		fmt.Println("----------4、私聊用户----------")
		fmt.Println("----------5、退出系统----------")
		fmt.Println("请选择(1-4):")
		var key int
		var content string
		smsMes := &SmsHandler{}

		fmt.Scanln(&key)
		switch key {
		case 1:
			//fmt.Println("显示在线用户列表")
			outPutOnlineUser()
		case 2:
			//fmt.Println("发送消息")
			fmt.Scanln(&content)
			smsMes.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Print("请输入私聊id：")
			var userId string
			fmt.Scanln(&userId)

		case 5:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入选项不正确(1-4)")
		}
	}
}

//和服务器保持通讯
func MessageHandler(conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	for {
		//fmt.Println("客户端正在等待读取服务器发射消息")
		mes, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("transfer.ReadPkg() err=", err)
			return
		}
		//fmt.Printf("mes=%v\n",mes)

		switch mes.Type {
		case common.NotifyUserStatusMesType:
			//有人上线了
			//取出notifyUserStatusMes
			var notifyUserStatusMes common.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			//把用户的信息，状态保存到客户map[string]User中
			updateUserStatus(&notifyUserStatusMes)
		case common.SmsMesType:
			outPutGroupMes(&mes)
		default:
			fmt.Println("服务器返回未知类型消息")
		}

	}
}
