package model

import (
	"encoding/json"
	"fmt"
	"github.com/TStarFall/golangIMSystem/common"
	"github.com/garyburd/redigo/redis"
)

//在服务器启动后，就初始化一个userDao实例
//定义一个全局变量的UserDao实例，需要操作redis时，直接使用
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}

	return
}

//1、根据用户id返回一个User实例
func (this *UserDao) getUserById(conn redis.Conn, id string) (user *User, err error) {
	//通过给定的id去redis查询这个用户
	res, err := redis.String(conn.Do("HGet","users", id))
	if err != nil {
		//从redis查询出错
		if err == redis.ErrNil{
			//标识在users中，没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	//实例化user
	user = &User{}
	//将查询到的res反序列化成user
	err = json.Unmarshal([]byte(res),user)
	if err != nil {
		fmt.Println("redis json.Unmarshal([]byte(res),&user) err=", err)
		return
	}
	return
}

//完成登录的校验逻辑
//1、Login完成对用户的验证
//2、如果用户的id 和 pwd都正确，则返回一个user实例
//3、如果用户id 和 pwd有错，则返回相对应的错误信息
func (this *UserDao) Login(userId, userPwd string) (user *User, err error) {
	//先从UserDao的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这时证明用户是可以获取到的
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *common.User) (err error) {
	//获取连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	//redis查询出错，说明redis中没有这个id存在
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	//写入redis
	_, err = conn.Do("HSet","users",user.UserId,string(data))
	if err != nil {
		fmt.Println("Register conn.Do err=", err)
		return
	}
	return
}






