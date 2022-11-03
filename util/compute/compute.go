package compute

import (
	"math/rand"
	"time"
)

// GetRandomString 获取随机字符串数字
func GetRandomString(len int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, len)

	for i := 0; i < len; i++ {
		random := rand.Intn(62)
		if random < 10 {
			result[i] = byte(48 + random)
		} else if random < 36 {
			result[i] = byte(55 + random)
		} else {
			result[i] = byte(61 + random)
		}
	}
	return string(result)
}
