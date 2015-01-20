package sabar

import (
	"github.com/jtolds/gls"
)

type Subject interface {
	GetId() string
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
	GetSession() *Session
	GetSessionId() string
	IsRemembered() bool
	//获取当前用户的所有权限信息
	GetPermissions() *HashSet
	GetRoles() *HashSet
}

//默认实现的subject
type DefaultSubject struct {
	name          string   //登陆名
	authenticated bool     //是否已经授权登陆了
	remembered    bool     //是否要记住
	roles         *HashSet //角色集合
	permissions   *HashSet //权限信息
	session       *Session
}

func (this *DefaultSubject) GetId() string {
	return this.GetId()
}
func (this *DefaultSubject) Login(token AuthcToken) UnAuthErr {
	info, err := realRealm.Authorize(token)
	if err != nil {
		return err
	}
	this.name = info.GetName()
	this.authenticated = true
	this.remembered = token.RememberMe
	this.roles = info.GetRoles()

}

type SubjectManager struct {
	lock     sync.Mutex
	Expires  int //过期时间
	subjects map[string]*Subject
}

var subjectManger *SubjectManager = &SubjectManager{
	Expires:  3600,
	subjects: make(map[string]*Subject),
}

//缓存所有的变量
//subjects := make(map[string]Subject)

//获取当前的Subject,未登陆时，生成
func GetSubject(id string) (*Subject, string) {
	if id == "" {
		subject := &DefaultSubject{}
		subject.id = idWorker.HexId()
		subject.session, id = sessionManager.NewSession()
		return subject, subject.id
	}
	return subjectManger.subjects[id], id
}
