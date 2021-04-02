package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
)

func outPutGroupMes(mes *common.Message) {
	//显示
	//1、反序列化mes.Data
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outPutGroupMes json.Unmarshal err=",err)
		return
	}

	info := fmt.Sprintf("用户id:%v 对大家说:%v",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
}
