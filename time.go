package godash

import (
	"errors"
	"time"
)

// https://www.cnblogs.com/jkko123/p/6909368.html

// GetCurrentTime 获取当前时间
func (g Godash) GetCurrentTime() time.Time {
	return time.Now().In(g.TimeLocation)
}

// GetCurrentDayBegin 获取当前时间的0点
func (g Godash) GetCurrentDayBegin() time.Time {
	now := g.GetCurrentTime()
	year, month, day := now.Date()
	return time.Date(year, month, day, 00, 00, 00, 00, g.TimeLocation)
}

// DateEqual 比较两个时间是否是同一天
func (g Godash) DateEqual(day1, day2 time.Time) bool {
	// 放在同一个时区
	day1 = day1.In(g.TimeLocation)
	day2 = day2.In(g.TimeLocation)

	// 转换成年月日
	y1, m1, d1 := day1.Date()
	y2, m2, d2 := day2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// Time2TimeStampMill 获取指定时间的毫秒级别的时间戳
func (g Godash) Time2TimeStampMill(t time.Time) int64 {
	return t.UnixNano() / 1e6 // 1e6 = 1000000
}

// 时间戳转时间

// Second2Time 秒 -> time.Time
func (g Godash) Second2Time(s int64) time.Time {
	return time.Unix(s, 0).In(g.TimeLocation)
}

// Millisecond2Time 毫秒 -> time.Time
func (g Godash) Millisecond2Time(s int64) time.Time {
	// 1 毫秒=1000000 纳秒
	return time.Unix(0, s*1e6).In(g.TimeLocation)
}

// Microsecond2Time 微秒 -> time.Time
func (g Godash) Microsecond2Time(s int64) time.Time {
	//1 微秒 = 1000 纳秒
	return time.Unix(0, s*1e3).In(g.TimeLocation)
}

// Nanosecond2Time 纳秒 -> time.Time
func (g Godash) Nanosecond2Time(s int64) time.Time {
	return time.Unix(0, s).In(g.TimeLocation)
}


const (
	MYNano      = "2006-01-02 15:04:05.000000000"
	MYMicro     = "2006-01-02 15:04:05.000000"
	MYMil       = "2006-01-02 15:04:05.000"
	MYSec       = "2006-01-02 15:04:05"
	MYCST       = "2006-01-02 15:04:05 +0800 CST"
	MYUTC       = "2006-01-02 15:04:05 +0000 UTC"
	MYDate      = "2006-01-02"
	MYTime      = "15:04:05"
	FBTIME      = "2006-01-02T15:04:05+0800"
	APPTIME     = "2006-01-02T15:04:05.000"
	TWITTERTIME = "2006-01-02T15:04:05Z"
)

// TimeStr2Time 字符串格式时间 转 time.Time
func (g Godash) TimeStr2Time(timeStr string) (time.Time, error) {
	// 可能的转换格式
	useFormat := []string{
		MYNano, MYMicro, MYMil, MYSec, MYCST, MYUTC, MYDate, MYTime, FBTIME, APPTIME, TWITTERTIME,
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
	}
	var t time.Time
	for _, useF := range useFormat {
		tt, err1 := time.ParseInLocation(useF, timeStr, g.TimeLocation)
		if err1 != nil {
			continue
		}
		t = tt
		break
	}
	// 没有找到任何符合时间格式
	if t == g.getTimeDefault() { // 0001-01-01 00:00:00 +0000 UTC
		return t, errors.New("时间字符串格式错误")
	}
	return t, nil
}

func (g Godash) getTimeDefault() time.Time {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "", g.TimeLocation)
	return t
}
