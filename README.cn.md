# godash 

[English](https://github.com/duanshanghanqing/godash/README.md) | 简体中文

## 安装
    go get -u github.com/duanshanghanqing/godash

## 使用
    import (
        "fmt"
        "time"
        "github.com/duanshanghanqing/godash"
    )

    type TimeSource struct {
        Name      string
        Age       int
        Is        bool
        Time      time.Time
    }

    type StringDestination struct {
        Name      string
        Age       int
        Is        bool
        Time      string
    }

    type TimeDestination struct {
        Name      string
        Age       int
        Is        bool
        Time      time.Time
    }
    
    func main() {
        tt := time.Now()
        var timeSource = TimeSource{
            Name:      "John",
            Age:       30,
            Is:        true,
            Time:      tt,
        }

        var stringDestination StringDestination
        godash.MappingStruct(&timeSource, &stringDestination)
        fmt.Println(stringDestination)

        var stringDestinationFormat StringDestination
        godash.MappingStructOption(&timeSource, &stringDestinationFormat, godash.Option{ FormatLayout: time.DateTime })
        fmt.Println(stringDestinationFormat)

        var timeDestination TimeDestination
        godash.MappingStruct(&timeSource, &timeDestination)
        fmt.Println(timeDestination)
    }

# 函数介绍
## MappingStruct
> 映射时间类型支持

    *time.Time -> string
    *time.Time -> *time.Time
    *time.Time -> time.Time

	time.Time -> string
    time.Time -> *time.Time
    time.Time -> time.Time

	string -> string
    string -> *time.Time
    string -> time.Time

## GetFreePort
> 获取动态端口

## GetIPv4
> 获取 IPv4 地址