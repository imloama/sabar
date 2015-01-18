package sabar

import (
	. "base"
	"reflect"
)

type Realm interface {
	//get name
	GetName() string
	//校验信息
	//Authenticate(token AuthcToken) (AuthcInfo, error)
	//从其他途径获取的用户登陆信息（如：数据库，其中可能包含的是加密后的用户信息）
	GetAuthcInfo(token AuthcToken) (AuthcInfo, error)
	//根据条件，获取验证完成用户的相关授权信息
	GetAuthzInfo(principal interface{}) AuthzInfo
}

//带加密处理策略的realm
type AuthRealm interface {
	Realm
	//密码加密处理策略
	Crypto
}

//参数类型的realm
type PropsRealm struct {
	Name        string
	Password    string
	Roles       []string
	Permissions []string
}

func (self *PropsRealm) GetName() string {
	return "propsRealm"
}

func (self *PropsRealm) GetAuthcInfo(token AuthcToken) (AuthcInfo, error) {
	return DefaultAuthcInfo{
		Name:     self.Name,
		Password: self.Password,
	}, nil
}

func (self *PropsRealm) GetAuthzInfo(principal interface{}) AuthzInfo {
	if principal == null {
		return nil
	}
	info := defaultAuthzInfo{
		name: self.Name,
	}
	if self.Roles != nil {
		info.roles = NewHashSet(self.Roles...)
	}
	if self.Permissions != nil {
		info.permissions = newHashSet(self.Permissions)
	}
	return info
}
