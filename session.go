package sabar

type Session interface {
	GetId() string
	SetAttribute(key, value interface{})
	GetAttribute(key interface{})
	RemoveAttribute(key interface{})
}
