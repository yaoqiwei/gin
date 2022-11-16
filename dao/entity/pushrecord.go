package entity

import "time"

type Pushrecord struct {
	Id        int64     `json:"id"`
	Genre     byte      `json:"genre"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Thumb     string    `json:"thumb"`
	Adminid   int64     `json:"adminid"`
	Admin     string    `json:"admin"`
	Ip        int64     `json:"ip"`
	Status    byte      `json:"status"`
	Disabled  byte      `json:"disabled"`
	Addtime   int64     `json:"addtime"`
	Pushtime  int64     `json:"pushtime"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
