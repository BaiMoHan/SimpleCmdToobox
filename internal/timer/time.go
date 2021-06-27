package timer

import (
	"log"
	"time"
)

// 获取当前时间
func GetNowTime() time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal("err=%v",err)
	}
	return time.Now().In(location)
}

// 推算时间
func GetCalculateTime(currentTimer time.Time,d string)(time.Time,error)  {
	// ParseDuration解析一个时间段字符串。一个时间段字符串是一个序列，
	// 每个片段包含可选的正负号、十进制数、可选的小数部分和单位后缀，
	// 如"300ms"、"-1.5h"、"2h45m"。
	// 合法的单位有"ns"、"us" /"µs"、"ms"、"s"、"m"、"h"。
	duration,err := time.ParseDuration(d)
	if err != nil{
		return time.Time{}, err
	}
	// 当前时间+duration
	return currentTimer.Add(duration),nil
}
