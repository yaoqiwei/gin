package pushrecord

import (
	"fmt"
	"gin/common/lib/gorm"
	"gin/dao/entity"
	"gin/model/body"
)

func Add(p body.PushrecordParam) {
	gorm.GinDb.Table("pushrecord").Create(&p)
}

func Get(id int64) {
	var pushrecord entity.Pushrecord
	err := gorm.GinDb.Where("id=?", id).FirstOrInit(&pushrecord).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("aaa", pushrecord)
}
