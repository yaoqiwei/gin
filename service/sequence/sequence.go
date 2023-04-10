package sequence

import "gin/dao/mapper/sequence"

// Top 置顶
func Top(id int64, tableName string, fileName string) {
	sequence.Top(id, tableName, fileName)
}

// MoveUp 上移
func MoveUp(id int64, tableName string, fileName string) {
	sequence.Move(id, tableName, fileName, 1)
}

// MoveDown
func MoveDown(id int64, tableName string, fileName string) {
	sequence.Move(id, tableName, fileName, 2)
}
