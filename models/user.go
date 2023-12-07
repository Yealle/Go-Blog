package models

import "time"

type User struct {
	Uid      int       `json:"json"`
	UserName string    `json:"userName"`
	Passwd   string    `json:"passwd"`
	Avatar   string    `json:"avatar"`
	CreateAt time.Time `json:"crerate_at"`
	UpdateAt time.Time `json:"updateAt_at"`
}

type UserInfo struct {
	Uid      int    `json:"json"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
