package handler

import "fmt"

//因为UserMgr实例在服务器端只有一个
//因为在很多地方都要用，因此定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[string]*UserConn
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[string]*UserConn,1024),
	}
}

func (this *UserMgr) AddOnlineUser(user *UserConn) {
	this.onlineUsers[user.UserId] = user
}

func (this *UserMgr) DelOnlineUser(userId string) {
	delete(this.onlineUsers, userId)
}

func (this *UserMgr) GetAllOnlineUser() map[string]*UserConn {
	return this.onlineUsers
}

func (this *UserMgr) GetOnlineUserById(userId string)(user *UserConn, err error) {
	user, ok := this.onlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%s 不存在",userId)
		return
	}
	return
}



