//相关异常封装
package sabar

import (
	"reflect"
)

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

func (this Err) Error() string {
	return fmt.Sprintf("Error:%d %s", this.Code, this.Msg)
}

func NewUnAuthErr(code int, msg string) UnAuthErr {
	return UnAuthErr{
		Code: code,
		Msg:  msg,
	}
}

func NewUnkownAccountErr(code int, msg string) UnkownAccountErr {
	return UnkownAccountErr{
		Code: code,
		Msg:  msg,
	}
}

func IsUnAuthErr(err error) bool {
	t := reflect.TypeOf(err)
	if t == UnAuthErr {
		return true
	}
	return false
}

func IsUnkownAccountErr(err error) bool {
	t := reflect.TypeOf(err)
	if t == UnkownAccountErr {
		return true
	}
	return false
}
