package compute

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

// RoundedFixed 保留小数点后n位-四舍五入
func RoundedFixed(val float64, n int) float64 {
	// 10的n次方
	point := math.Pow10(n)
	fv := 0.0000000001 + val //对浮点数产生.xxx999999999 计算不准进行处理
	return math.Floor(fv*point+.5) / point
}

// TruncRound 舍去小数点后n位
func TruncRound(val float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n+1)+"f", val)
	temp := strings.Split(floatStr, ".")
	var newFloat string
	if len(temp) < 2 || n > len(temp[1]) {
		newFloat = floatStr
	} else {
		newFloat = temp[0] + "." + temp[1][:n]
	}
	inst, _ := strconv.ParseFloat(newFloat, 64)
	return inst
}

// GetRandomInt 获取随机字符串数字
func GetRandomInt(len int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, len)
	for i := 0; i < len; i++ {
		result[i] = byte(48 + rand.Intn(10))
	}
	return string(result)
}

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

// Rand returns a random
func Rand(min, max int) int {
	if min > max {
		panic("min: min cannot be greater than max")
	}
	if int31 := 1<<31 - 1; max > int31 {
		panic("max: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// AddFloat 浮点数相加
func AddFloat(a, b float64) float64 {
	res, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Float64()
	return res
}

// SubFloat 浮点数相减
func SubFloat(a, b float64) float64 {
	res, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Float64()
	return res
}

// DivFloat 浮点数相除
func DivFloat(a, b float64) float64 {
	res, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Float64()
	return res
}

// MulFloat 浮点数相乘
func MulFloat(a, b float64) float64 {
	res, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Float64()
	return res
}

// Min 返回最小值
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Max 返回最大值
func Max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
