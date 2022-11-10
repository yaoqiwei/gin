package entity

import "time"

// Users 用户表结构
type Users struct {
	Id             int64     `json:"id"`
	UserLogin      string    `json:"userLogin"`      // 账号
	UserNickname   string    `json:"userNickName"`   // 用户美名
	Avatar         string    `json:"avatar"`         // 头像
	AvatarThumb    string    `json:"avatarThumb"`    // 头像缩略图
	RealName       string    `json:"realName"`       // 真实姓名
	UserPass       string    `json:"-"`              // 用户密码
	Mobile         string    `json:"mobile"`         // 用户手机号
	UserEmail      string    `json:"userEmail"`      // 用户邮箱
	Sex            int8      `json:"sex"`            // 用户性别: 0:保密 1:男 2:女
	Birthday       string    `json:"birthday"`       // 生日
	Signature      string    `json:"signature"`      // 用户个性签名
	Location       string    `json:"location"`       // 所在地
	PrivateKey     string    `json:"-"`              // 密码密钥
	HardwareId     string    `json:"hardwareId"`     // 硬件id
	UserStatus     int8      `json:"userStatus"`     // 用户状态 0：禁用； 1：正常
	UserType       byte      `json:"userType"`       //0 普通用户，1 admin
	LastLoginIp    string    `json:"lastLoginIp"`    // 最后登录ip
	LastLoginTime  time.Time `json:"lastLoginTime"`  // 最后登录时间
	Token          string    `json:"-"`              // 授权token
	ExpireTime     int64     `json:"-"`              // token到期时间
	ValidationFlag int64     `json:"validationFlag"` // 是否二次校验 0:否，1：是
}
