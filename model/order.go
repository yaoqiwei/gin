package model

import (
	"fmt"
	"gin/model/body"
	"time"
)

type OrderData struct {
	Num int64  `json:"num"`
	Day string `json:"day"`
}

func OrderSuccess(p body.OrderNumParam) {
	now := time.Now()
	fmt.Println(now)
	list := make([]OrderData, 0)
	err := DB.Table("cmf_order").
		Select("count(*) num,DATE_FORMAT(create_at,'%Y-%m-%d %H:00:00') day").Debug().
		Where("create_at between ? and ?", p.StartTime, p.EndTime).
		Group("day").Order("day desc").
		Scan(&list).Error
	fmt.Println("list", list, err)
}
