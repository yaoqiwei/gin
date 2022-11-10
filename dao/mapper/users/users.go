package users

import (
	"gin/common/lib/gorm"
	"gin/dao/entity"
	"gin/model/body"
)

// Get 获取用户详情
func Get(p body.UserSearchParam) (entity.Users, error) {
	users := entity.Users{}
	err := gorm.GinDb.DB.Model(&entity.Users{}).
		Where("id =?", p.Uid).Limit(1).
		Find(&users).Error
	return users, err
}
