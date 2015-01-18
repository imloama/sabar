//相关认证
package sabar

//要校验的信息
type AuthcInfo interface {
	//获取当登陆名
	GetPrincipal() interface{}
	//获取登陆密码
	GetCredentials() interface{}
	//密码盐
	GetSalt() interface{}
}

type DefaultAuthcInfo struct {
	Name     string //登陆名
	Password string //密码
	Salt     string //密码盐
}

func (self *DefaultAuthcInfo) GetPrincipal() interface{} {
	return self.Name
}

func (self *DefaultAuthcInfo) GetCredentials() interface{} {
	return self.Password
}
func (self *DefaultAuthcInfo) GetSalt() interface{} {
	return self.Salt
}
