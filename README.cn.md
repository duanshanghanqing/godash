# godash 

[English](https://github.com/duanshanghanqing/godash/README.md) | 简体中文

一个go语言工具库

## 安装使用
    go get -u github.com/duanshanghanqing/godash

## 初始化一个实例
    import (
        "fmt"
        "github.com/duanshanghanqing/godash"
    )

    func main() {
        // 使用默认配置
        //config := godash.Config{
        //	LoadLocationName: "Asia/Shanghai",
        //	Layout: "2006-01-02 15:04:05",
        //}
        g := godash.Default()
        
        // 或指定配置
        //g := godash.New(config)
    }

## 时间处理

        // 获取当前时间
        GetCurrentTime() time.Time

        // 获取当前时间的0点
        GetCurrentDayBegin() time.Time

        // 比较两个时间是否是同一天
        DateEqual(day1, day2 time.Time) bool

        // 获取指定时间的毫秒级别的时间戳
        Time2TimeStampMill(t time.Time) int64

        // 秒 转 time.Time
        Second2Time(s int64) time.Time

        // 毫秒 转 time.Time
        Millisecond2Time(s int64) time.Time

        // 微秒 转 time.Time
        Microsecond2Time(s int64) time.Time
        
        // 纳秒 转 time.Time
        Nanosecond2Time(s int64) time.Time

        // 字符串格式时间 转 time.Time
        TimeStr2Time(timeStr string) (time.Time, error)

## 实例映射
    
####支持16种转换

	*time.Time -> int64
	*time.Time -> time.Time
	*time.Time -> *time.Time
	*time.Time -> string

	time.Time -> int64
	time.Time -> *time.Time
	time.Time -> time.Time
	time.Time -> string

	int64 -> int64
	int64 -> *time.Time
	int64 -> time.Time
    int64 -> string

	string -> int64
	string -> *time.Time
	string -> time.Time
	string -> string

####例子

    g := Default()

	type Stu struct {
		Str  string
		Time *time.Time
	}
	type Stu2 struct {
		Str  string
		Time int64
		Name string
	}

	_time := time.Date(2021, 10, 01, 23, 45, 44, 00, g.TimeLocation)
	stu := Stu{Str: "test", Time: &_time}
	stu2 := Stu2{}

	g.EntityMapping(stu, &stu2)
	fmt.Println(stu2)