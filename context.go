//上下文变量
package sabar

import (
	. "github.com/itwarcraft/idgenerator"
)

var sessionManager *SessionManager = NewSessionManager()

//实际realm
var realRealm *Realm

idWorker, _ := NewIdWorker(0, 0)

//设置实际的realm
func InitRealm(realm *Realm) {
	realRealm = realm
}
