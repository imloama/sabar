package sabar

import (
    "github.com/jtolds/gls"
)

type Subject interface {
    GetId()string
	//登陆
	Login(token AuthcToken) UnAuthErr
	//退出
	Logout()
	//获取登陆名
	GetName() string
	//是否被授权
	IsPermitted(permission string) bool
	//是否已验证
	IsAuthenticated() bool
	//GetSession() Session
	IsRemembered() bool
	//获取当前用户的所有权限信息
	GetPermissions() *HashSet
	GetRoles() *HashSet
}

//默认实现的subject
type DefaultSubject struct {
	Name          string   //登陆名
	password      []byte   //密码
	authenticated bool     //是否已经授权登陆了
	remembered    bool     //是否要记住
	permissions   []string //权限信息
}

func (self *DefaultSubject) Login(token AuthcToken) UnAuthErr {

}
//缓存所有的变量
subjects := make(map[string]Subject)


//获取当前的Subject,未登陆时，生成
func GetSubject() *Subject {
	return nil
}
