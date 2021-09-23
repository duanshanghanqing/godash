# godash

English | [简体中文](https://github.com/duanshanghanqing/godash/blob/master/README.cn.md)

a go language tool library

## Installation and use
    go get -u github.com/duanshanghanqing/godash

## Initialize an instance
    import (
        "fmt"
        "github.com/duanshanghanqing/godash"
    )

    func main() {
        // Use default configuration
        //config := godash.Config{
        //	LoadLocationName: "Asia/Shanghai",
        //	Layout: "2006-01-02 15:04:05",
        //}
        g := godash.Default()
        
        // Or specify the configuration
        //g := godash.New(config)
    }

## Time processing

        // Get current time
        GetCurrentTime() time.Time

        // Get 0 point of current time
        GetCurrentDayBegin() time.Time

        // Compare whether the two times are the same day
        DateEqual(day1, day2 time.Time) bool

        // Gets the timestamp at the millisecond level of the specified time
        Time2TimeStampMill(t time.Time) int64

        // Second to time.time
        Second2Time(s int64) time.Time

        // millisecond to time.time
        Millisecond2Time(s int64) time.Time

        // Microsecond to time.Time
        Microsecond2Time(s int64) time.Time
        
        // Nanosecond to time.Time
        Nanosecond2Time(s int64) time.Time

        // String format time to time.Time
        TimeStr2Time(timeStr string) (time.Time, error)

## Instance mapping

#### Support 16 kinds of conversion

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

#### Example

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