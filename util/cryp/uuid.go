package cryp

import "github.com/google/uuid"

// GetUUID 获取唯一id
func GetUUID() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
