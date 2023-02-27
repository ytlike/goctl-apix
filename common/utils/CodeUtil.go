package utils

import (
	"qbq-open-platform/common/global"
	"strconv"
	"time"
)

var tH = time.Now().Hour()
var tMi = time.Now().Minute()
var tS = time.Now().Second()
var tNaS = time.Now().Nanosecond() / 1e6

func GetFormatTime(time time.Time) string {
	return time.Format("20060102")
}

func GetTime(date string) string {
	return date + strconv.Itoa(tH) + strconv.Itoa(tMi) + strconv.Itoa(tS)
}

func GetCode(prefix string) string {
	date := GetFormatTime(time.Now())
	//获取redis增长序号
	number, err := global.Config().RedisClient.Incr(global.CACHE_CODE_UTIL_VALUE)
	if err != nil {
		return ""
	}
	var code = strconv.FormatInt(number, 10)
	length := 8 - len(code)
	//补足8位
	if length > 0 {
		for i := 0; i < length; i++ {
			code = "0" + code
		}
	}
	return prefix + GetTime(date) + code
}
