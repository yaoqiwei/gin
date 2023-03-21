package pushrecord

import (
	"gin/dao/mapper/pushrecord"
	"gin/model/body"
)

// PushrecordAdd 添加消息
func PushrecordAdd(p body.PushrecordParam) {
	pushrecord.Get(7)
	// pushrecord.Add(p)
}
