package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
	"github.com/TStarFall/golangIMSystem/server/model"
	"github.com/TStarFall/golangIMSystem/server/utils"
	"net"
)

type UserConn struct {
	Conn net.Conn

	//增加UserId字段
	UserId string
}

func (this *UserConn) RegisterHandler(mes *common.Message) (err error) {
	//1、从mes中取出mes.Data，直接反序列化成RegisterMes
	var registerMes common.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("RegisterHandler json.Unmarshal err=",err)
		return
	}
	//声明resMes，用于存放回传信息
	var resMes common.Message
	resMes.Type = common.RegisterResMesType
 	var registerResMes common.RegisterResMes
	//到redis完成注册
	//1、使用model.MyUserDao到数据库验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册时发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}

	//registerResMes序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("RegisterHandler json.Marshal(registerResMes) err=",err)
		return
	}
	//data 赋值给resMes.Data
	resMes.Data = string(data)

	//resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("RegisterHandler json.Marshal(resMes) err=",err)
		return
	}

	//回传resMes给客户端
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

func (this *UserConn)LoginHandler(mes *common.Message) (err error) {
	//核心代码
	//1、先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal mes.Data err=", err)
		return err
	}
	//fmt.Println("loginMes=", loginMes)
	//2、声明一个reMes
	var resMes common.Message
	resMes.Type = common.LoginResMesType
	//3、声明一个loginResMes
	var loginResMes common.LoginResMes
	//到redis进行查询验证
	user, err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		//将登陆成功的用户的userId 赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其它在线用户，我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前在线用户的id放入到loginResMes.UserId
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserId = append(loginResMes.UserId, id)
		}

		fmt.Println(user,"登陆成功")
	}

	//4、将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err=", err)
		return err
	}

	//5、将data赋值给resMes
	resMes.Data = string(data)

	//6、resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err=", err)
		return err
	}
	//发送消息
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("writePkg(conn,data) err=", err)
		return err
	}
	return
}

//通知所有在线用户的方法
//userId 要通知其他的在线用户，我上线了
func (this *UserConn) NotifyOthersOnlineUser(userId string) {
	//遍历onlineUsers，然后一个一个的发送 NotifyUsersStatusMes
	for id, user := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}
		user.NotifyMeOnline(userId)
	}
}

func (this *UserConn) NotifyMeOnline(userId string) {
	//
	var mes common.Message
	mes.Type = common.NotifyUserStatusMesType

	var notifyUserStatusMes common.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = common.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal(notifyUserStatusMes) err=", err)
		return
	}

	//将序列化后的notifyUserStatusMes赋给mes.data
	mes.Data = string(data)

	//mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("NotifyMeOnline json.Marshal(mes) err=", err)
		return
	}

	//发送mes
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline transfer.WritePkg(data) err=", err)
		return
	}
}


