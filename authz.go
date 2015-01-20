//校验动作相关处理功能
package sabar

import (
	. "base"
)

//校验通过后的信息存储
type AuthzInfo interface {
	//获取当前登陆名
	GetName() string
	//当前验证用户的角色
	GetRoles() *HashSet
	//当前验证用户的权限
	GetPermissions() *HashSet
	//其他信息存储
	GetAttrs() map[interface{}]interface{}
}

//通用性校验后的信息存储
type defaultAuthzInfo struct {
	name        String
	roles       *HashSet
	permissions *HashSet
	attrs       map[interface{}]interface{}
}

//返回验证通过的
func (self *defaultAuthzInfo) GetName() string {
	return self.name
}

func (self *defaultAuthzInfo) GetRoles() *HashSet {
	return self.roles
}
func (self *defaultAuthzInfo) GetPermissions() *HashSet {
	return self.permissions
}
func (self *defaultAuthzInfo) GetAttrs() map[interface{}]interface{} {
	return self.attrs
}
