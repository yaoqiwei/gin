package convert

import (
	"encoding/json"
	"sort"
)

// ToJson
func ToJson(i interface{}, e ...string) string {
	b, err := json.Marshal(i)
	if err != nil {
		if len(e) > 0 {
			return e[0]
		}
		return ""
	}
	return string(b)
}

// InArrayString 判断数组中是否存在该值
func InArrayString(s string, arr []string) bool {
	for _, i := range arr {
		if i == s {
			return true
		}
	}
	return false
}

// In 判断数组中是否存在该值
func In(s string, arr []string) bool {
	sort.Strings(arr)
	index := sort.SearchStrings(arr, s)
	if index < len(arr) && arr[index] == s {
		return true
	}
	return false
}
