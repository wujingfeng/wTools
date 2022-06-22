package util

import (
	"time"
)

const sec_format = "2006-01-02 15:04:05"
const day_format = "2006-01-02"

// Today
//  @Description: 今天的日期
//  @return string
//  @author	wujingfeng 2022-06-22 11:54:49
func Today() string {
	n := time.Now()
	return n.Format(day_format)
}

// TodayLastSeconds
//  @Description: 今日剩余秒数
//  @return int
//  @author	wujingfeng 2022-06-22 11:54:58
func TodayLastSeconds() int64 {
	// 获取明天的日期
	tomorrowSec := DaysSeconds(1)
	tomorrowDate := time.Unix(int64(tomorrowSec), 0).Format(day_format)
	// 明天开始的时间戳
	tomorrowStartSec := DateToTimestamp(tomorrowDate)
	// 返回剩余秒数
	return tomorrowStartSec - time.Now().Unix()
}

// DateToTimestamp
//  @Description: 日期转时间戳
//  @param date string eg: "2022-06-22"
//  @return int64
//  @author	wujingfeng 2022-06-22 14:07:36
func DateToTimestamp(date string) int64 {
	t, err := time.ParseInLocation(day_format, date, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// DatetimeToTimestamp
//  @Description: 日期时间转时间戳
//  @param date	string  eg: "2022-06-22 14:20:20"
//  @return int64
//  @author	wujingfeng 2022-06-22 14:08:32
func DatetimeToTimestamp(date string) int64 {
	t, err := time.ParseInLocation(sec_format, date, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// DaysSeconds
//  @Description: 指定天数后的时间戳
//  @param days
//  @return int64
//  @author	wujingfeng 2022-06-22 12:01:18
func DaysSeconds(days int64) int64 {
	t := time.Now().Unix()
	t = t + days*86400
	return time.Unix(t, 0).Unix()
}
