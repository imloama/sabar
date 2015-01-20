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
	//验证功能
	Authorize(AuthcToken token) (AuthzInfo, error)
}

//参数类型的realm
type PropsRealm struct {
	name        string
	password    string
	roles       []string
	permissions []string
}

func (this *PropsRealm) GetName() string {
	return "propsRealm"
}

func (this *PropsRealm) GetAuthcInfo(token AuthcToken) (AuthcInfo, error) {
	return DefaultAuthcInfo{
		name:     this.name,
		password: this.password,
	}, nil
}

func (this *PropsRealm) GetAuthzInfo(principal interface{}) AuthzInfo {
	if principal == null {
		return nil
	}
	info := defaultAuthzInfo{
		name: self.name,
	}
	if this.roles != nil {
		info.roles = NewHashSet(this.roles...)
	}
	if this.permissions != nil {
		info.permissions = newHashSet(this.permissions)
	}
	return info
}

func (this *PropsRealm) Authorize(AuthcToken token) (AuthzInfo, error) {
	if token.GetPrincipal() == this.Name && token.GetCredentials() == this.Password {
		return this.GetAuthzInfo(), nil
	}
	return nil, UnAuthErr{
		Code: -1,
		Msg:  "name or password error!",
	}
}

//带加密处理策略的realm
type AuthRealm interface {
	Realm
	//密码加密处理策略
	Crypto
}
