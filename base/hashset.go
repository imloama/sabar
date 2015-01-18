//hashset实现，实现来源于 go并发编程,来源 https://github.com/hyper-carrot/goc2p/blob/master/src/basic/set/hash_set.go
package base

import (
	"bytes"
	"fmt"
)

type Set interface {
	Add(val interface{}) bool
	Len() int
	Remove(val interface{})
	Clear()
	Contains(val interface{}) bool
	Elements() []interface{}
	String() string
}

//HashSet实现
type HashSet struct {
	attrs map[interface{}]bool
}

//添加
func (self *HashSet) Add(val interface{}) bool {
	if !self.attrs[val] {
		self.attrs[val] = true
		return true
	}
	return false
}

func (self *HashSet) Len(i int) interface{} {
	return len(self.attrs)
}

func (self *HashSet) Remove(val interface{}) {
	delete(self.attrs, val)
}

func (self *HashSet) Clear() {
	self.attrs = make(map[interface{}]bool)
}

func (self *HashSet) Contains(val interface{}) bool {
	return self.attrs[val]
}

func (self *HashSet) Elements() []interface{} {
	initLen := len(self.attrs)
	cp := make([]interface{}, initLen)
	actLen := 0
	for key := range self.attrs {
		if actLen < initLen {
			cp[actLen] = key
		} else {
			cp = append(cp, key)
		}
		actLen++
	}
	if actLen < initLen {
		cp = cp[:actLen]
	}
	return cps
}
func (self *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("HashSet{")
	first := true
	for key := range self.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func NewHashSet(val ...interface{}) *HashSet {
	set := &HaseSet{
		attrs: make(map[interface{}]bool),
	}
	if val != nil {
		for _, v := range val {
			set.Add(v)
		}
	}
	return set
}
