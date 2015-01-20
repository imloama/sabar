//session实现，内容参考 http://blog.guoqiangti.com/?p=318
package sabar

import (
	"container/list"
	"log"
	"sync"
	"time"
)

type Session interface {
	GetId() string
	SetAttr(key, value interface{}) *Session
	GetAttr(key interface{}) interface{}
	RemoveAttr(key interface{})
	//获取过期时间
	GetExpires() int
}

//默认session实现，为内存模式
type DefaultSession struct {
	lock    sync.Mutex
	id      string
	attrs   map[interface{}]interface{}
	Expires int //过期时间
}

//session管理器
type SessionManager struct {
	lock     sync.Mutex
	Expires  int //过期时间
	sessions map[string]*Session
}

func (this *DefaultSession) GetId() string {
	return this.id
}

func (this *DefaultSession) SetAttr(key, value interface{}) *Session {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.attrs[key] = value
	return this
}

func (this *DefaultSession) GetAttr(key interface{}) interface{} {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this[key]
}

func (this *DefaultSession) RemoveAttr(key interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.attrs, key)
}

func NewSessionManager() *SessionManager {
	sm := SessionManager{}
	sm.sessions = make(map[string]*Session)
	sm.Expires = 3600
	return &sm
}

func (this *SessionManager) NewSession() (*Session, string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	id := idWorker.HexId()
	session := DefaultSession{
		id:      id,
		attrs:   make(map[interface{}]interface{}),
		Expires: 3600,
	}
	this.sessions[id] = &session
	return &session, id
}

func (this *SessionManager) GetSession(id string) *Session {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.sessions[id]
}
