package model

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功，我们必须保证
	//用户信息的json字符串的key和结构体字段对应的tag一致
	UserId string `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}

