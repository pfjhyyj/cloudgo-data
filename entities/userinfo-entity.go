package entities

import (
	"time"
)

// UserInfo .
type UserInfo struct {
	UID        int
	UserName   string
	DepartName string
	CreateAt   *time.Time `xorm:"created"`
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}
