package util

import "encoding/json"

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
