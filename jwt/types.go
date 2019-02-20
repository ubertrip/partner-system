package jwt

import (
	"unsafe"
)

type Session struct {
	Login   bool   `json:"login,omitempty"`
	ID      int64  `json:"id,omitempty"`
	Expired int64  `json:"exp,omitempty"`
	GroupID uint32 `json:"grp,omitempty"`
}

func (s *Session) String() (str string) {
	data, _ := json.Marshal(s)
	return *(*string)(unsafe.Pointer(&data))
}
