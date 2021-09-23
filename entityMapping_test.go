package godash

import (
	"fmt"
	"testing"
	"time"
)

/*
	支持16种转换
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
	time.Time -> int64

	string -> int64
	string -> *time.Time
	string -> time.Time
	string -> string
*/

// *time.Time -> int64
// go test -v -run Test_EntityMapping_Time2Int64 entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMapping_Time2Int64(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time *time.Time
		Height int64
		Width int64
	}
	type Stu2 struct {
		Str  string
		Time int64
		Name string
	}
	_time := time.Date(2021, 10, 01, 23, 45, 44, 00, g.TimeLocation)
	stu := Stu{
		Str: "test",
		Time: &_time,
		Height: 176,
		Width: 100,
	}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu2)
}

// *time.Time -> time.Time
// go test -v -run Test_EntityMapping_Time2Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMapping_Time2Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time *time.Time
	}
	type Stu2 struct {
		Str  string
		Time time.Time
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: &_time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// *time.Time -> *time.Time
// go test -v -run Test_EntityMapping_Time2_Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMapping_Time2_Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time *time.Time
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: &_time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// *time.Time -> string
// go test -v -run Test_EntityMapping_Time2String entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMapping_Time2String(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time *time.Time
	}
	type Stu2 struct {
		Str  string
		Time string
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: &_time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// time.Time -> int64
// go test -v -run Test_EntityMappingTime2int64 entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingTime2int64(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time time.Time
	}
	type Stu2 struct {
		Str  string
		Time int64
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: _time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// time.Time -> *time.Time
// go test -v -run Test_EntityMapping_Time2Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingTime2_Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time time.Time
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: _time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// time.Time -> time.Time
// go test -v -run Test_EntityMappingTime2Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingTime2Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time time.Time
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: _time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// time.Time -> string
// go test -v -run Test_EntityMappingTime2String entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingTime2String(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time time.Time
	}
	type Stu2 struct {
		Str  string
		Time string
		Name string
	}
	var _time = time.Now()
	stu := Stu{Str: "test", Time: _time}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//int64 -> int64
// go test -v -run Test_EntityMappingInt642int64 entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingInt642int64(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time int64
	}
	type Stu2 struct {
		Str  string
		Time int64
		Name string
	}
	stu := Stu{Str: "test", Time: int64(1633103144000)}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//int64 -> *time.Time
// go test -v -run Test_EntityMappingInt642_Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingInt642_Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time int64
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	stu := Stu{Str: "test", Time: int64(1633103144000)}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//int64 -> time.Time
// go test -v -run Test_EntityMappingInt642Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingInt642Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time int64
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	stu := Stu{Str: "test", Time: int64(1633103144000)}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//int64 -> string
// go test -v -run Test_EntityMappingInt642String entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingInt642String(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time int64
	}
	type Stu2 struct {
		Str  string
		Time string
		Name string
	}
	stu := Stu{Str: "test", Time: int64(1633103144000)}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

// 2021-10-01 23:45:44 -> int64
// go test -v -run Test_EntityMappingStr2Int64 entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingStr2Int64(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time string
	}
	type Stu2 struct {
		Str  string
		Time int64
		Name string
	}
	stu := Stu{Str: "test", Time: "2021-10-01 23:45:44"}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//2021-10-01 23:45:44 -> *time.Time
//go test -v -run Test_EntityMappingStr2Int64 entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingStr2_Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time string
	}
	type Stu2 struct {
		Str  string
		Time *time.Time
		Name string
	}
	stu := Stu{Str: "test", Time: "2021-10-01 23:45:44"}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//2021-10-01 23:45:44 -> time.Time
// go test -v -run Test_EntityMappingStr2Time entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingStr2Time(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time string
	}
	type Stu2 struct {
		Str  string
		Time time.Time
		Name string
	}
	stu := Stu{Str: "test", Time: "2021-10-01 23:45:44"}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}

//2021-10-01 23:45:44 -> time.Time
// go test -v -run Test_EntityMappingStr2Str entityMapping_test.go entityMapping.go time.go godash.go
func Test_EntityMappingStr2Str(t *testing.T) {
	g := Default()
	type Stu struct {
		Str  string
		Time string
	}
	type Stu2 struct {
		Str  string
		Time string
		Name string
	}
	stu := Stu{Str: "test", Time: "2021-10-01 23:45:44"}
	stu2 := Stu2{}
	g.EntityMapping(stu, &stu2)
	fmt.Println(stu)
	fmt.Println(stu2)
}
