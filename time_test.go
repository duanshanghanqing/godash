package godash

import (
	"fmt"
	"testing"
	"time"
)

// go test -v -run TestGetCurrentTime time_test.go  time.go godash.go
func TestGetCurrentTime(t *testing.T) {
	//g := Default()
	g := New(Config{
		LoadLocationName: "Asia/Shanghai",
	})
	now := g.GetCurrentTime()
	fmt.Println(now)
}

// go test -v -run TestGetCurrentDayBegin time_test.go  time.go godash.go
func TestGetCurrentDayBegin(t *testing.T) {
	g := Default()
	now := g.GetCurrentDayBegin()
	fmt.Println(now)
}

// go test -v -run TestDateEqual time_test.go  time.go godash.go
func TestDateEqual(t *testing.T) {
	g := Default()
	time1 := time.Date(2021, 10, 01, 23, 45, 44, 00, g.TimeLocation)
	time2 := time.Date(2021, 10, 01, 23, 45, 44, 00, g.TimeLocation)
	fmt.Println(g.DateEqual(time1, time2))
}

// go test -v -run TestTime2TimeStampMill time_test.go  time.go godash.go
func TestTime2TimeStampMill(t *testing.T) {
	g := Default()
	time := time.Date(2021, 10, 01, 23, 45, 44, 00, g.TimeLocation)
	fmt.Println(g.Time2TimeStampMill(time))
}

// go test -v -run TestSecond2Time time_test.go  time.go godash.go
func TestSecond2Time(t *testing.T) {
	g := Default()
	fmt.Println(g.Second2Time(1632298377))
}

// go test -v -run TestMillisecond2Time time_test.go  time.go godash.go
func TestMillisecond2Time(t *testing.T) {
	g := Default()
	fmt.Println(g.Millisecond2Time(1632298780981))
}

// go test -v -run TestSubtle2Time time_test.go  time.go godash.go
func TestSubtle2Time(t *testing.T) {
	g := Default()
	fmt.Println(g.Microsecond2Time(1.6323e+15))
}

// go test -v -run TestNanosecond2Time time_test.go  time.go godash.go
func TestNanosecond2Time(t *testing.T) {
	g := Default()
	fmt.Println(g.Nanosecond2Time(1.6323e+18))
}

// go test -v -run TestTimeStr2Time time_test.go  time.go godash.go
func TestTimeStr2Time(t *testing.T) {
	g := Default()
	_time1, err1 := g.TimeStr2Time("2021-09-22 17:15:34")
	if err1 != nil {
		return
	}
	fmt.Println("_time1", _time1)

	_time2, err2 := g.TimeStr2Time("2021-09-22")
	if err2 != nil {
		return
	}
	fmt.Println("_time2", _time2)
}
