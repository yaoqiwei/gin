package sequence

import (
	"fmt"
	"gin/common/lib/gorm"

	gorm2 "gorm.io/gorm"
)

type TableNo struct {
	Id int64
	No int64
}

// Top 置顶
func Top(id int64, tableName, fileName string) {
	err := gorm.GinDb.Transaction(func(tx *gorm2.DB) error {
		if err := tx.Table(tableName).Updates(map[string]interface{}{
			fileName: gorm2.Expr(fileName+"+?", 1),
		}).Error; err != nil {
			return err
		}
		return tx.Model(tableName).Where("id = ?", id).Update(fileName, 1).Error
	})
	if err != nil {
		panic(err.Error())
	}
}

// Move 移动
func Move(id int64, tableName, fileName string, genre byte) {
	var currentNo int64
	var err error

	err = gorm.GinDb.Transaction(func(tx *gorm2.DB) error {
		if err = tx.Table(tableName).
			Where("id = ?", id).
			Pluck(fileName, &currentNo).Error; err != nil {
			return err
		}

		tableNo := TableNo{}
		if genre == 1 {
			// 上移
			tableNo, err = MaxCurrentNo(tx, tableName, fileName, currentNo)
			if err != nil {
				return err
			}
			if tableNo.Id == 0 {
				return nil
			}

		} else {
			// 下移
			tableNo, err = MinCurrentNo(tx, tableName, fileName, currentNo)
			if err != nil {
				return err
			}

			if tableNo.Id == 0 {
				return nil
			}
		}

		if err = tx.Table(tableName).Where("id = ?", tableNo.Id).Update(fileName, currentNo).Error; err != nil {
			return err
		}

		return tx.Table(tableName).Where("id = ?", tableNo.Id).Update(fileName, tableNo.No).Error

	})

	if err != nil {
		panic(err.Error())
	}
}

// MaxCurrentNo 获取当前位置上方的最大值
func MaxCurrentNo(db *gorm2.DB, tableName, fileName string, currentNo int64) (TableNo, error) {
	tableNo := TableNo{}
	sql := fmt.Sprintf("select id,max(%s) no from %s where %s < %d", fileName, tableName, fileName, currentNo)
	if err := db.Raw(sql).Limit(1).Find(&tableNo).Error; err != nil {
		return tableNo, err
	}

	err := db.Table(tableName).Select("id").Where(fileName+" = ? ", tableNo.No).
		Pluck("id", &tableNo.Id).Error
	return tableNo, err
}

// MinCurrentNo 获取当前位置下方的最小值
func MinCurrentNo(db *gorm2.DB, tableName, fileName string, currentNo int64) (TableNo, error) {
	tableNo := TableNo{}
	sql := fmt.Sprintf("select id,min(%s) no from %s where %s > %d", fileName, tableName, fileName, currentNo)
	if err := db.Raw(sql).Limit(1).Find(&tableNo).Error; err != nil {
		return tableNo, err
	}
	err := db.Table(tableName).Select("id").Where(fileName+" = ? ", tableNo.No).
		Pluck("id", &tableNo.Id).Error
	return tableNo, err
}
