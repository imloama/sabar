//相关异常封装
package sabar

//异常的封装
type Err struct {
	Code int    //异常代码
	Msg  string //异常提示信息
}

//未授权用户
type UnAuthErr struct {
	Err
}

//未知用户
type UnkownAccountErr struct {
	Err
}

func (self Err) Error() string {
	return fmt.Sprintf("Error:%d %s", self.Code, self.Msg)
}
