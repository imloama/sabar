//加密相关
//部分加密算法内容来源于：https://github.com/williambao/gocrypto/blob/master/crypto.go
package sabar

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
)

//加密接口
type Crypto interface {
	//加密
	//Encrypt(source []byte) ([]byte, error)
	//获取当前的信息，完成密码加密，并生成新的authcinfo对象
	Encrypt(info *AuthcInfo) (*AuthcInfo, error)
}

//默认加密处理
type DefaultCrypto struct {
}

//密码加密默认处理方法
func (self *DefaultCrypto) Encrypt(info *AuthcInfo) (*AuthcInfo, error) {
	c0 := GetMD5(info.Password)
	c := GetHmacSha1(c, info.Salt)
	return &DefaultAuthcInfo{
		Name:     info.Name,
		Password: c,
		Salt:     info.Salt,
	}, nil
}

func GetSha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func GetHmacSha1(data string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func GetMD5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckMD5(data string, sum string) bool {
	out := GetMD5(data)
	return sum == out
}

func Base64EncodeByte(data []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(data))
}

func Base64DecodeByte(data []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(data))
}

func Base64Encode(data string) string {
	return string(Base64EncodeByte([]byte(data)))
}

func Base64Decode(data string) (string, error) {
	d, e := Base64DecodeByte([]byte(data))
	return string(d), e
}
