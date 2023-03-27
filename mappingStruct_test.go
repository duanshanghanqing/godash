package godash

import (
	"testing"
	"time"
)

type User struct {
	Height int64
}

type PointerTimeSource struct {
	Name      string
	Age       int
	Is        bool
	Time      *time.Time
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

type TimeSource struct {
	Name      string
	Age       int
	Is        bool
	Time      time.Time
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

type StringTimeSource struct {
	Name      string
	Age       int
	Is        bool
	Time      string
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

var pointerTimeSource PointerTimeSource
var timeSource TimeSource
var stringTimeSource StringTimeSource

type StringDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      string
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

type TimeDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      time.Time
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

type PointerTimeDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      *time.Time
	TestMap   map[string]string
	TestSlice []string
	Arr       [10]float32
	User
	User1     User
	User2     *User
	MyChannel chan User
	GetUser   func()
	TypeName  interface{}
	Price     string
}

func init() {
	//time.Local = time.FixedZone("UTC", 0)

	testMap := make(map[string]string)
	testMap["name"] = "berners"
	testSlice := make([]string, 1)
	testSlice = append(testSlice, "a")
	myChannel := make(chan User, 1)
	myChannel <- User{Height: 120}
	tt := time.Now()

	pointerTimeSource = PointerTimeSource{
		Name:      "John",
		Age:       30,
		Is:        true,
		Time:      &tt,
		TestMap:   testMap,
		User:      User{Height: 100},
		User1:     User{Height: 110},
		User2:     &User{Height: 120},
		MyChannel: myChannel,
		GetUser:   func() {},
		TypeName:  1,
	}

	timeSource = TimeSource{
		Name:      "John",
		Age:       30,
		Is:        true,
		Time:      tt,
		TestMap:   testMap,
		User:      User{Height: 100},
		User1:     User{Height: 110},
		User2:     &User{Height: 120},
		MyChannel: myChannel,
		GetUser:   func() {},
		TypeName:  1,
	}

	stringTimeSource = StringTimeSource{
		Name:      "John",
		Age:       30,
		Is:        true,
		Time:      "2023-03-23T14:39:14.8562572+08:00",
		TestMap:   testMap,
		User:      User{Height: 100},
		User1:     User{Height: 110},
		User2:     &User{Height: 120},
		MyChannel: myChannel,
		GetUser:   func() {},
		TypeName:  1,
	}
}

// go test -v mappingStruct_test.go mappingStruct.go

// *time.Time -> string
// *time.Time -> *time.Time
// *time.Time -> time.Time

// go test -v -run Test_mapStruct_pointerTimeToString mappingStruct_test.go mappingStruct.go
func Test_mapStruct_pointerTimeToString(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 StringDestination
	MappingStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MappingStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_pointerTimeToPointerTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_pointerTimeToPointerTime(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 PointerTimeDestination
	MappingStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MappingStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_pointerTimeToTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_pointerTimeToTime(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 TimeDestination
	MappingStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MappingStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// ----------------------------------------------------------------------------------------

// time.Time -> string
// time.Time -> *time.Time
// time.Time -> time.Time

// go test -v -run Test_mapStruct_timeToString mappingStruct_test.go mappingStruct.go
func Test_mapStruct_timeToString(t *testing.T) {
	t.Log(timeSource)

	var dest1 StringDestination
	MappingStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MappingStruct(timeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_timeToPointerTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_timeToPointerTime(t *testing.T) {
	t.Log(timeSource)

	var dest1 PointerTimeDestination
	MappingStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MappingStruct(timeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_timeToTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_timeToTime(t *testing.T) {
	t.Log(timeSource)

	var dest1 TimeDestination
	MappingStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MappingStruct(timeSource, &dest2)
	t.Log(dest2)
}

// ----------------------------------------------------------------------------------------

// string -> string
// string -> *time.Time
// string -> time.Time

// go test -v -run Test_mapStruct_stringTimeToString mappingStruct_test.go mappingStruct.go
func Test_mapStruct_stringTimeToString(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 StringDestination
	MappingStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MappingStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_stringTimeToPointerTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_stringTimeToPointerTime(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 PointerTimeDestination
	MappingStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MappingStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_stringTimeToTime mappingStruct_test.go mappingStruct.go
func Test_mapStruct_stringTimeToTime(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 TimeDestination
	MappingStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MappingStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}

// ----------------------------------------------------------------------------------------

type GoodsSource struct {
	Uint   string
	Uint8  string
	Uint16 string
}

type GoodsDestination struct {
	Uint   uint
	Uint8  uint8
	Uint16 uint16
}

// go test -v -run Test_mapStruct_stringToInt mappingStruct_test.go mappingStruct.go
func Test_mapStruct_stringToInt(t *testing.T) {
	var goodsSource GoodsSource
	goodsSource.Uint = "1"
	goodsSource.Uint8 = "2"
	goodsSource.Uint16 = "3"

	var goodsDestination GoodsDestination
	MappingStruct(goodsSource, &goodsDestination)
	t.Log(goodsDestination)
}
