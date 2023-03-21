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

func Get(id int64) error {
	var pushrecord entity.Pushrecord
	// err := gorm.GinDb.Debug().FirstOrInit(&pushrecord, id).Error
	var title string
	gorm.GinDb.Model(&pushrecord).Select("title").Where(id).Debug().Row().Scan(&title)
	fmt.Println(pushrecord, title)
	// err := gorm.GinDb.Where("id=?", id).FirstOrInit(&pushrecord).Error
	// return err
	return nil
}
