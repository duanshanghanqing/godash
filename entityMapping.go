package godash

import (
	"reflect"
	"time"
)

// EntityMapping p_接收的是实体，f_ 接收的是指针
func (g Godash) EntityMapping(t interface{}, p interface{}) {
	getType, getValue := reflect.TypeOf(t), reflect.ValueOf(t)
	_, setValue := reflect.TypeOf(p), reflect.ValueOf(p)
	v := setValue.Elem()

	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)

		// 左边的 字段 不在右边，就跳过
		if v.FieldByName(field.Name).String() == "<invalid Value>" {
			continue
		}


		// 左侧类型
		leftTimeType := field.Type.String()
		// 右侧类型
		rightTimeType := v.FieldByName(field.Name).Type().String()


		// time.Time 转 *time.Time 或 time.Time
		var timeToTime = func(_time interface{}) {
			_, getTimeValue := reflect.TypeOf(_time), reflect.ValueOf(_time)
			v.FieldByName(field.Name).Set(getTimeValue)
		}

		switch leftTimeType {
		// 左侧是 *time.Time 类型
		case "*time.Time":
			// 把 *time.Time 类型，转换为 int64 类型, 毫秒级别
			int64Time := getValue.Field(i).Interface().(*time.Time).UnixNano() / 1e6
			// 右侧接收的 int64 类型。毫秒级别
			if rightTimeType == "int64" {
				v.FieldByName(field.Name).SetInt(int64Time)
			} else if rightTimeType == "time.Time" {
				// 毫秒转 time
				_time := g.Millisecond2Time(int64Time)
				timeToTime(_time)
			} else if rightTimeType == "*time.Time" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				v.FieldByName(field.Name).SetString(getValue.Field(i).Interface().(*time.Time).Format(g.Layout))
			}
			break
		case "time.Time":
			int64Time := getValue.Field(i).Interface().(time.Time).UnixNano() / 1e6
			if rightTimeType == "int64" {
				v.FieldByName(field.Name).SetInt(int64Time)
			} else if rightTimeType == "*time.Time" {
				// 毫秒转 time
				_time := g.Millisecond2Time(int64Time)
				timeToTime(&_time)
			} else if rightTimeType == "time.Time" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				v.FieldByName(field.Name).SetString(getValue.Field(i).Interface().(time.Time).Format(g.Layout))
			}
			break
		case "int64": // 毫秒级别 js new Date().getTime()返回值
			if rightTimeType == "*time.Time" {
				// 前端 getTime()返回值 毫秒 ， 转 time。 毫秒 -> time
				_time := g.Millisecond2Time(getValue.Field(i).Int())
				timeToTime(&_time)
			} else if rightTimeType == "time.Time" {
				// 前端 getTime()返回值 毫秒 ， 转 time。 毫秒 -> time
				_time := g.Millisecond2Time(getValue.Field(i).Int())
				timeToTime(_time)
			} else if rightTimeType == "int64" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				// 前端 getTime()返回值 毫秒 ， 转 time。 毫秒 -> time
				_time := g.Millisecond2Time(getValue.Field(i).Int())
				v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
			}
			break
		case "string": // 字符串时间转换成 int64 毫秒级别， *time.Time， time.Time
			_time, _timeErr := g.TimeStr2Time(getValue.Field(i).String())
			// 说明 字符串日期转换成功
			if _timeErr == nil {
				if rightTimeType == "int64" {
					// 转为毫秒
					int64Time := _time.UnixNano() / 1e6
					v.FieldByName(field.Name).SetInt(int64Time)
				} else if rightTimeType == "*time.Time" {
					timeToTime(&_time)
				} else if rightTimeType == "time.Time" {
					timeToTime(_time)
				} else if rightTimeType == "string" {
					v.FieldByName(field.Name).Set(getValue.Field(i))
				}
			} else {
				v.FieldByName(field.Name).Set(getValue.Field(i))
			}
			break
		default:
			v.FieldByName(field.Name).Set(getValue.Field(i))
			break
		}
	}
}
