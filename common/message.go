package common

//定义两个信息类型常量
const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"`  //消息类型
	Data string  `json:"data"`//消息的类型
}

//定义登陆消息
type LoginMes struct {
	UserId string `json:"userId"` //用户Id
	UserPwd string `json:"userPwd"`//用户Pwd
	UserName string `json:"userName"`//用户名
}

//定义登陆消息返回
type LoginResMes struct {
	Code int  `json:"code"`//返回状态码
	Error string `json:"error"`//返回状态信息
	UserId []string `json:"userId"`
}

type RegisterMes struct {
	//...
	User User `json:"user"`
}

type RegisterResMes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}

//为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId string `json:"userId"`
	Status int `json:"status"`
}

type SmsMes struct {
	Content string `json:"content"`//发送的内容
	User //匿名结构体
}
