package godash

import (
	"errors"
	"reflect"
	"time"
)

func TimeStrToTime(timeStr string) (time.Time, error) {
	layouts := []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,
	}
	var t = time.Now()
	var ts = t.String()

	for _, layout := range layouts {
		tt, err := time.Parse(layout, timeStr)
		if err != nil {
			continue
		}
		t = tt
		break
	}
	// No matching time format found
	if t.String() == ts {
		return t, errors.New("datetime string format error")
	}
	return t, nil
}

func mappingStruct(l interface{}, r interface{}, opt Option) {
	getType, getValue := reflect.TypeOf(l), reflect.ValueOf(l)
	_, setValue := reflect.TypeOf(r), reflect.ValueOf(r)

	// right check Pointer
	if setValue.Kind() != reflect.Ptr {
		return
	}
	if setValue.IsNil() {
		return
	}

	if getValue.Kind() == reflect.Ptr && !getValue.IsNil() {
		v := setValue.Elem()

		for i := 0; i < getType.Elem().NumField(); i++ {
			// 左字段
			leftField := getType.Elem().Field(i)
			// 左值
			leftValue := getValue.Elem().Field(i)

			// 左边的值空指针
			if leftValue.Kind() == reflect.Ptr && leftValue.IsNil() {
				//fmt.Println("左侧空指针：", field.Name, "---", getValue.Elem().Field(i))
				continue
			}

			// 左边的 字段 不在右边，就跳过
			rightField := v.FieldByName(leftField.Name)
			if rightField.String() == "<invalid Value>" {
				//fmt.Println("左边的字段不在右边：", field.Name, "---", getValue.Elem().Field(i))
				continue
			}

			// 左侧类型
			leftFieldType := leftField.Type.String()
			// 右侧类型
			rightFieldType := rightField.Type().String()

			switch leftFieldType {
			case "*time.Time":
				// *time.Time -> string
				// *time.Time -> *time.Time
				// *time.Time -> time.Time

				_time := leftValue.Interface().(*time.Time)
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).SetString(_time.Format(opt.FormatLayout))
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(leftValue)
				} else if rightFieldType == "time.Time" {
					// *time.Time -> string -> time.Time
					t, _ := TimeStrToTime(_time.Format(opt.FormatLayout))
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(t))
				}
				break
			case "time.Time":
				// time.Time -> string
				// time.Time -> *time.Time
				// time.Time -> time.Time

				_time := leftValue.Interface().(time.Time)
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).SetString(_time.Format(opt.FormatLayout))
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(&_time))
				} else if rightFieldType == "time.Time" {
					v.FieldByName(leftField.Name).Set(leftValue)
				}
				break
			case "string":
				// string -> string
				// string -> *time.Time
				// string -> time.Time
				_time, _ := TimeStrToTime(leftValue.String())
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).Set(leftValue)
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(&_time))
				} else if rightFieldType == "time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(_time))
				}
				break
			default:
				// 两侧数据类型一至才可以映射
				if leftFieldType == rightFieldType &&
					leftValue.Kind().String() != "ptr" &&
					leftValue.Kind().String() != "array" &&
					leftValue.Kind().String() != "struct" &&
					leftValue.Kind().String() != "chan" &&
					leftValue.Kind().String() != "func" &&
					leftValue.Kind().String() != "slice" &&
					leftValue.Kind().String() != "interface" &&
					leftValue.Kind().String() != "map" {
					v.FieldByName(leftField.Name).Set(leftValue)
				}
				break
			}
		}
	}

	if getValue.Kind() == reflect.Struct {
		v := setValue.Elem()
		for i := 0; i < getType.NumField(); i++ {
			// 左字段
			leftField := getType.Field(i)
			// 左值
			leftValue := getValue.Field(i)

			// 左边的值空指针
			if leftValue.Kind() == reflect.Ptr && leftValue.IsNil() {
				//fmt.Println("左侧空指针：", field.Name, "---", getValue.Field(i))
				continue
			}
			// 左边的 字段 不在右边，就跳过
			if v.FieldByName(leftField.Name).String() == "<invalid Value>" {
				//fmt.Println("左边的字段不在右边：", field.Name, "---", getValue.Field(i))
				continue
			}

			// 左侧类型
			leftFieldType := leftField.Type.String()
			// 右侧类型
			rightFieldType := v.FieldByName(leftField.Name).Type().String()

			switch leftFieldType {
			// 左侧是 *time.Time 类型
			case "*time.Time":
				// *time.Time -> string
				// *time.Time -> *time.Time
				// *time.Time -> time.Time

				_time := leftValue.Interface().(*time.Time)
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).SetString(_time.Format(opt.FormatLayout))
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(leftValue)
				} else if rightFieldType == "time.Time" {
					// *time.Time -> string -> time.Time
					t, _ := TimeStrToTime(_time.Format(opt.FormatLayout))
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(t))
				}
				break
			// 左侧是 time.Time 类型
			case "time.Time":
				// time.Time -> string
				// time.Time -> *time.Time
				// time.Time -> time.Time

				_time := leftValue.Interface().(time.Time)
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).SetString(_time.Format(opt.FormatLayout))
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(&_time))
				} else if rightFieldType == "time.Time" {
					v.FieldByName(leftField.Name).Set(leftValue)
				}
				break
			// 左侧是 time.Time 类型
			case "string":
				// string -> string
				// string -> *time.Time
				// string -> time.Time

				_time, _ := TimeStrToTime(leftValue.String())
				if rightFieldType == "string" {
					v.FieldByName(leftField.Name).Set(leftValue)
				} else if rightFieldType == "*time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(&_time))
				} else if rightFieldType == "time.Time" {
					v.FieldByName(leftField.Name).Set(reflect.ValueOf(_time))
				}
				break
			default:
				// 两侧数据类型一至才可以映射
				if leftFieldType == rightFieldType &&
					leftValue.Kind().String() != "ptr" &&
					leftValue.Kind().String() != "array" &&
					leftValue.Kind().String() != "struct" &&
					leftValue.Kind().String() != "chan" &&
					leftValue.Kind().String() != "func" &&
					leftValue.Kind().String() != "slice" &&
					leftValue.Kind().String() != "interface" &&
					leftValue.Kind().String() != "map" {
					v.FieldByName(leftField.Name).Set(leftValue)
				}
				break
			}
		}
	}
}

type Option struct {
	FormatLayout string
}

func MappingStruct(l interface{}, r interface{}) {
	mappingStruct(l, r, Option{FormatLayout: time.RFC3339Nano})
}

func MappingStructOption(l interface{}, r interface{}, opt Option) {
	mappingStruct(l, r, opt)
}
