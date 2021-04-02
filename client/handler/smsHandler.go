package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/client/utils"
	"github.com/TStarFall/golangIMSystem/common"
)

type SmsHandler struct {

}

//发送群聊的消息
func (this *SmsHandler) SendGroupMes(content string) (err error){
	//创建一个mes
	var mes common.Message
	mes.Type = common.SmsMesType

	//创建一个SmsMes 实例
	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化
	data, err := json.Marshal(smsMes)
	if err !=nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) err=",err)
		return
	}

	//mes.Data赋值
	mes.Data = string(data)

	//mes 序列化
	data, err = json.Marshal(mes)
	if err !=nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err=",err)
		return
	}

	//发送消息
	transfer := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err !=nil {
		fmt.Println("SendGroupMes transfer.WritePkg(data) err=",err)
		return
	}
	return
}

func (this *SmsHandler) SendToOneUserMes(userId, content string) (err error) {
	var mes common.Message
	mes.Type = common.SmsMesType

	var smsMes common.SmsMes

	user, ok := onlineUsers[userId]
	if !ok {
		fmt.Println("该用户不在线...")
		return
	}

	smsMes.Content = content
	smsMes.User.UserId = user.UserId
	smsMes.User.UserStatus = user.UserStatus

	//序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendToOneUserMes json.Marshal(smsMes) err=",err)
		return
	}

	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendToOneUserMes json.Marshal(mes) err=",err)
		return
	}

	transfer := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("SendToOneUserMes transfer.WritePkg err=",err)
		return
	}
	return
}
