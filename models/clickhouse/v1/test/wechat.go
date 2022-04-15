package test

import (
	"gin/models"
	"time"
)

type Wechatusersessions struct {
	ID      uint64
	Seq     uint64
	Msgid   string
	Msgtype string
	Msgtime time.Time
	// Msgtime    models.GormTime
	Action     string
	Roomid     string
	From       string
	Created_at models.GormTime
	Updated_at models.GormTime
}

func (Wechatusersessions) TableName() string {
	return "wechat_user_sessions"
}
