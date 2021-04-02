package handler

import (
	"fmt"
	"github.com/TStarFall/golangIMSystem/client/model"
	"github.com/TStarFall/golangIMSystem/common"
)

//客户端要维护的map
var onlineUsers map[string]*common.User = make(map[string]*common.User,10)
var CurUser model.CurUser //用户登陆成功后，完成对CurUser的初始化
//在客户端显示当前在线用户
func outPutOnlineUser() {
	fmt.Println("在线用户列表：")
	//遍历onlineUsers
	for id, _ := range onlineUsers {
		fmt.Printf("用户id:%v\n",id)
	}
}

func updateUserStatus(notifyUserStatusMes *common.NotifyUserStatusMes){
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &common.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}

	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outPutOnlineUser()
}

