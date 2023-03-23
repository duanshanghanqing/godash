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
        godash.MapStruct(&timeSource, &stringDestination)
        fmt.Println(stringDestination)

        var stringDestinationFormat StringDestination
        godash.MapStructOption(&timeSource, &stringDestinationFormat, godash.Option{ FormatLayout: time.DateTime })
        fmt.Println(stringDestinationFormat)

        var timeDestination TimeDestination
        godash.MapStruct(&timeSource, &timeDestination)
        fmt.Println(timeDestination)
    }

# 函数介绍
## MapStruct
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