package godash

import "time"

type Config struct {
	LoadLocationName string
	Layout           string
}

type Godash struct {
	TimeLocation *time.Location
	Layout       string
}

func New(config Config) Godash {
	g := Godash{}
	// 初始化时区
	timeLocation, err := time.LoadLocation(config.LoadLocationName)
	if err != nil {
		panic(err)
	}
	g.TimeLocation = timeLocation
	g.Layout = config.Layout
	return g
}

func Default() Godash {
	config := Config{
		LoadLocationName: "Asia/Shanghai",
		Layout:           "2006-01-02 15:04:05",
	}
	return New(config)
}
