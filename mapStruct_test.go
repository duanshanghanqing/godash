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
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
}

type TimeSource struct {
	Name      string
	Age       int
	Is        bool
	Time      time.Time
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
}

type StringTimeSource struct {
	Name      string
	Age       int
	Is        bool
	Time      string
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
}

var pointerTimeSource PointerTimeSource
var timeSource TimeSource
var stringTimeSource StringTimeSource

type StringDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      string
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
}

type TimeDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      time.Time
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
}

type PointerTimeDestination struct {
	Name      string
	Age       int
	Is        bool
	Time      *time.Time
	testMap   map[string]string
	testSlice []string
	arr       [10]float32
	User
	user      User
	user2     *User
	myChannel chan User
	getUser   func()
	typeName  interface{}
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
		testMap:   testMap,
		User:      User{Height: 100},
		user:      User{Height: 110},
		user2:     &User{Height: 120},
		myChannel: myChannel,
		getUser:   func() {},
		typeName:  1,
	}

	timeSource = TimeSource{
		Name:      "John",
		Age:       30,
		Is:        true,
		Time:      tt,
		testMap:   testMap,
		User:      User{Height: 100},
		user:      User{Height: 110},
		user2:     &User{Height: 120},
		myChannel: myChannel,
		getUser:   func() {},
		typeName:  1,
	}

	stringTimeSource = StringTimeSource{
		Name:      "John",
		Age:       30,
		Is:        true,
		Time:      "2023-03-23T14:39:14.8562572+08:00",
		testMap:   testMap,
		User:      User{Height: 100},
		user:      User{Height: 110},
		user2:     &User{Height: 120},
		myChannel: myChannel,
		getUser:   func() {},
		typeName:  1,
	}
}

// go test -v mapStruct_test.go mapStruct.go

// *time.Time -> string
// *time.Time -> *time.Time
// *time.Time -> time.Time

// go test -v -run Test_mapStruct_pointerTimeToString mapStruct_test.go mapStruct.go
func Test_mapStruct_pointerTimeToString(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 StringDestination
	MapStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MapStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_pointerTimeToPointerTime mapStruct_test.go mapStruct.go
func Test_mapStruct_pointerTimeToPointerTime(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 PointerTimeDestination
	MapStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MapStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_pointerTimeToTime mapStruct_test.go mapStruct.go
func Test_mapStruct_pointerTimeToTime(t *testing.T) {
	t.Log(pointerTimeSource)

	var dest1 TimeDestination
	MapStruct(&pointerTimeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MapStruct(pointerTimeSource, &dest2)
	t.Log(dest2)
}

// ----------------------------------------------------------------------------------------

// time.Time -> string
// time.Time -> *time.Time
// time.Time -> time.Time

// go test -v -run Test_mapStruct_timeToString mapStruct_test.go mapStruct.go
func Test_mapStruct_timeToString(t *testing.T) {
	t.Log(timeSource)

	var dest1 StringDestination
	MapStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MapStruct(timeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_timeToPointerTime mapStruct_test.go mapStruct.go
func Test_mapStruct_timeToPointerTime(t *testing.T) {
	t.Log(timeSource)

	var dest1 PointerTimeDestination
	MapStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MapStruct(timeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_timeToTime mapStruct_test.go mapStruct.go
func Test_mapStruct_timeToTime(t *testing.T) {
	t.Log(timeSource)

	var dest1 TimeDestination
	MapStruct(&timeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MapStruct(timeSource, &dest2)
	t.Log(dest2)
}

// ----------------------------------------------------------------------------------------

// string -> string
// string -> *time.Time
// string -> time.Time

// go test -v -run Test_mapStruct_stringTimeToString mapStruct_test.go mapStruct.go
func Test_mapStruct_stringTimeToString(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 StringDestination
	MapStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 StringDestination
	MapStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_stringTimeToPointerTime mapStruct_test.go mapStruct.go
func Test_mapStruct_stringTimeToPointerTime(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 PointerTimeDestination
	MapStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 PointerTimeDestination
	MapStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}

// go test -v -run Test_mapStruct_stringTimeToTime mapStruct_test.go mapStruct.go
func Test_mapStruct_stringTimeToTime(t *testing.T) {
	t.Log(stringTimeSource)

	var dest1 TimeDestination
	MapStruct(&stringTimeSource, &dest1)
	t.Log(dest1)

	var dest2 TimeDestination
	MapStruct(stringTimeSource, &dest2)
	t.Log(dest2)
}
