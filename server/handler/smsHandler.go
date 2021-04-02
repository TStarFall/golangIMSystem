package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
	"github.com/TStarFall/golangIMSystem/server/utils"
	"net"
)

type SmsHandler struct {
	//暂时不写字段
}

//转发消息的方法
func (this *SmsHandler) SendGroupMes(mes *common.Message) {
	//遍历服务器端的onlineUsers map[string]*UserConn
	//将转发的消息取出
	//取出mes的内容SmsMes
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal([]byte(mes.Data), &smsMes) err=", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) err=", err)
		return
	}

	for _, user :=  range userMgr.onlineUsers {
		//if id == smsMes.UserId {
		//	continue
		//}
		this.SendMesToOtherOnlineUser(data, user.Conn)
	}
}

func (this *SmsHandler) SendMesToOtherOnlineUser(data []byte,conn net.Conn) {
	transfer := &utils.Transfer{
		Conn: conn,
	}
	err := transfer.WritePkg(data)
	if err != nil {
		fmt.Scanln("SendMesToOtherOnlineUser err=",err)
	}
}



