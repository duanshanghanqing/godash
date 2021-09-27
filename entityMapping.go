package godash

import (
	"fmt"
	"reflect"
	"time"
)

func (g Godash) EntityMapping(l interface{}, r interface{}) {
	getType, getValue := reflect.TypeOf(l), reflect.ValueOf(l)
	_, setValue := reflect.TypeOf(r), reflect.ValueOf(r)
	if setValue.Kind() == reflect.Ptr {
		if setValue.IsNil() {
			fmt.Println("the right cannot be a null pointer")
			return
		}
	} else {
		fmt.Println("the right must be a pointer")
		return
	}

	v := setValue.Elem()


	// 左侧指针类型
	if getValue.Kind() == reflect.Ptr  {
		// 左侧空指针，直接退出
		if getValue.IsNil() {
			fmt.Println("the left side cannot be a null pointer")
			return
		}
		for i := 0; i < getType.Elem().NumField(); i++ {
			field := getType.Elem().Field(i)

			// 左边的值空指针
			if getValue.Elem().Field(i).Kind() == reflect.Ptr && getValue.Elem().Field(i).IsNil() {
				//fmt.Println("左侧空指针：", field.Name, "---", getValue.Elem().Field(i))
				continue
			}
			// 左边的 字段 不在右边，就跳过
			if v.FieldByName(field.Name).String() == "<invalid Value>" {
				//fmt.Println("左边的字段不在右边：", field.Name, "---", getValue.Elem().Field(i))
				continue
			}

			// 左侧类型
			leftTimeType := field.Type.String()
			// 右侧类型
			rightTimeType := v.FieldByName(field.Name).Type().String()

			switch leftTimeType {
			// 左侧是 *time.Time 类型
			case "*time.Time":
				_time := getValue.Elem().Field(i).Interface().(*time.Time).In(g.TimeLocation)
				if rightTimeType == "int64" {
					// 时间戳方便前端解析，如果要更精确，可以使用纳秒
					v.FieldByName(field.Name).SetInt(_time.UnixNano() / 1e6)
				} else if rightTimeType == "time.Time" {
					v.FieldByName(field.Name).Set(reflect.ValueOf(g.Nanosecond2Time(_time.UnixNano())))
				} else if rightTimeType == "*time.Time" {
					// 两侧类型相同
					v.FieldByName(field.Name).Set(getValue.Elem().Field(i))
				} else if rightTimeType == "string" {
					v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
				}
				break
			case "time.Time":
				_time := getValue.Elem().Field(i).Interface().(time.Time).In(g.TimeLocation)
				// 把 *time.Time 类型，转换为 int64 类型, 毫秒级别
				int64Time := _time.UnixNano() / 1e6
				if rightTimeType == "int64" {
					v.FieldByName(field.Name).SetInt(int64Time)
				} else if rightTimeType == "*time.Time" {
					v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
				} else if rightTimeType == "time.Time" {
					// 两侧类型相同
					v.FieldByName(field.Name).Set(getValue.Elem().Field(i))
				} else if rightTimeType == "string" {
					v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
				}
				break
			case "int64": // 毫秒级别 js new Date().getTime()返回值
				// 前端 getTime()返回值 毫秒 ， 转 time。 毫秒 -> time
				_time := g.Millisecond2Time(getValue.Elem().Field(i).Int())
				if rightTimeType == "*time.Time" {
					v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
				} else if rightTimeType == "time.Time" {
					v.FieldByName(field.Name).Set(reflect.ValueOf(_time))
				} else if rightTimeType == "int64" {
					// 两侧类型相同
					v.FieldByName(field.Name).Set(getValue.Elem().Field(i))
				} else if rightTimeType == "string" {
					v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
				}
				break
			case "string": // 字符串时间转换成 int64 毫秒级别， *time.Time， time.Time
				if rightTimeType == "string" {
					v.FieldByName(field.Name).Set(getValue.Elem().Field(i))
				} else {
					// 说明 字符串日期转换成功
					_time, _timeErr := g.TimeStr2Time(getValue.Elem().Field(i).String())
					if _timeErr == nil {
						if rightTimeType == "int64" {
							// 转为毫秒
							int64Time := _time.UnixNano() / 1e6
							v.FieldByName(field.Name).SetInt(int64Time)
						} else if rightTimeType == "*time.Time" {
							v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
						} else if rightTimeType == "time.Time" {
							v.FieldByName(field.Name).Set(reflect.ValueOf(_time))
						}
					}
				}
				break
			default:
				// 两侧数据类型一至才可以映射
				if leftTimeType == rightTimeType {
					v.FieldByName(field.Name).Set(getValue.Elem().Field(i))
				}
				break
			}
		}
		return
	}



	// 左侧非指针类型
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)

		// 左边的值空指针
		if getValue.Field(i).Kind() == reflect.Ptr && getValue.Field(i).IsNil() {
			//fmt.Println("左侧空指针：", field.Name, "---", getValue.Field(i))
			continue
		}
		// 左边的 字段 不在右边，就跳过
		if v.FieldByName(field.Name).String() == "<invalid Value>" {
			//fmt.Println("左边的字段不在右边：", field.Name, "---", getValue.Field(i))
			continue
		}

		// 左侧类型
		leftTimeType := field.Type.String()
		// 右侧类型
		rightTimeType := v.FieldByName(field.Name).Type().String()

		switch leftTimeType {
		// 左侧是 *time.Time 类型
		case "*time.Time":
			_time := getValue.Field(i).Interface().(*time.Time).In(g.TimeLocation)
			// 右侧接收的 int64 类型。毫秒级别
			if rightTimeType == "int64" {
				// 把 *time.Time 类型，转换为 int64 类型, 毫秒级别
				v.FieldByName(field.Name).SetInt(_time.UnixNano() / 1e6)
			} else if rightTimeType == "time.Time" {
				v.FieldByName(field.Name).Set(reflect.ValueOf(g.Nanosecond2Time(_time.UnixNano())))
			} else if rightTimeType == "*time.Time" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
			}
			break
		case "time.Time":
			_time := getValue.Field(i).Interface().(time.Time).In(g.TimeLocation)
			if rightTimeType == "int64" {
				// 转毫秒输出
				v.FieldByName(field.Name).SetInt(_time.UnixNano() / 1e6)
			} else if rightTimeType == "*time.Time" {
				v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
			} else if rightTimeType == "time.Time" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
			}
			break
		case "int64": // 毫秒级别 js new Date().getTime()返回值
			// 前端 getTime()返回值 毫秒 ， 转 time。 毫秒 -> time
			_time := g.Millisecond2Time(getValue.Field(i).Int())
			if rightTimeType == "*time.Time" {
				v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
			} else if rightTimeType == "time.Time" {
				v.FieldByName(field.Name).Set(reflect.ValueOf(_time))
			} else if rightTimeType == "int64" {
				// 两侧类型相同
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else if rightTimeType == "string" {
				v.FieldByName(field.Name).SetString(_time.Format(g.Layout))
			}
			break
		case "string": // 字符串时间转换成 int64 毫秒级别， *time.Time， time.Time
			if rightTimeType == "string" {
				v.FieldByName(field.Name).Set(getValue.Field(i))
			} else {
				_time, _timeErr := g.TimeStr2Time(getValue.Field(i).String())
				// 说明 字符串日期转换成功
				if _timeErr == nil {
					if rightTimeType == "int64" {
						// 转为毫秒
						v.FieldByName(field.Name).SetInt(_time.UnixNano() / 1e6)
					} else if rightTimeType == "*time.Time" {
						v.FieldByName(field.Name).Set(reflect.ValueOf(&_time))
					} else if rightTimeType == "time.Time" {
						v.FieldByName(field.Name).Set(reflect.ValueOf(_time))
					}
				}
			}
			break
		default:
			// 两侧数据类型一至才可以映射
			if leftTimeType == rightTimeType {
				v.FieldByName(field.Name).Set(getValue.Field(i))
			}
			break
		}
	}

}
