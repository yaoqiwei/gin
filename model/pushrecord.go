package model

import (
	"time"
)

type Pushrecord struct {
	Id        int64             `json:"id" gorm:"primaryKey"`
	Touid     []InformPushTouid `json:"touid"`   // 推送对象
	Genre     byte              `json:"genre"`   // 公告类型 1-系统通知，2-活动通知
	Title     string            `json:"title"`   // 标题
	Content   string            `json:"content"` // 推送内容
	Thumb     string            `json:"thumb"`   // 内容图片链接
	Adminid   int64             `json:"-"`       // 管理员ID
	Admin     string            `json:"admin"`   // 管理员账号
	Ip        int64             `json:"-"`       // 管理员IP地址
	Status    byte              `json:"status"`  // 通知状态：1.草稿,2.待发布,3.已发布,4.已撤回,5.已下架
	Disabled  byte              `json:"-"`       // 0.显示，1.隐藏
	Addtime   int64             `json:"-"`       // 创建时间
	Pushtime  int64             `json:"-"`       // 推送时间
	CreatedAt time.Time
	UpdatedAt time.Time
}

// InformPushTouid 推送对象
type InformPushTouid struct {
	Id           int64  `json:"id"`           //用户id
	UserNicename string `json:"userNicename"` //用户名称
}
