//用户提供信息的封装
package sabar

//授权信息
type AuthcToken interface {
	//获取当登陆名
	GetPrincipal() interface{}
	//获取登陆密码
	GetCredentials() interface{}
	//获取salt
	GetSalt() interface{}
}

//用户名密码授权
type UserPwdToken struct {
	Name       string //登陆名
	ShowName   string //显示名称
	Password   string //密码
	RememberMe bool   //记住我
	Salt       string //加密盐
}

//存储新的 用户名密码token
func NewUserPwdToken(username string, password string, salt string) *UserPwdToken {
	return &UserPwdToken{
		Name:     username,
		Password: password,
		Salt:     salt,
	}
}

func (self *UserPwdToken) GetPrincipal() interface{} {
	return self.Name
}

func (self *UserPwdToken) GetCredentials() interface{} {
	return self.password
}

func (self *UserPwdToken) GetSalt() interface{} {
	return self.Salt
}
