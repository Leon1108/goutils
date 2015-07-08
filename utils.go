package goutils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"
	"unicode"
)

const (
	TIME_FOMRAT_SIMPLE = "2006-01-02 15:04:05.999"
)

// 获取当前的毫秒值
func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 获取格式化过的时间，精确到毫秒
func GetCurrentTime() string {
	return time.Now().Format(TIME_FOMRAT_SIMPLE)
}

// 判断当前运行方式是否为单元测试
// 仅适用于 `go test ...` 的执行方式
func IsTesting() bool {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	if strings.HasSuffix(path, "test") {
		return true
	}
	return false
}

// 将一个 interface{} 类型输出成字符串，主要是针对Struct类型进行了特殊处理
func ToString(obj interface{}) string {
	if obj == nil {
		return fmt.Sprintf("%v", obj)
	}

	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	// 识别输入参数类型
	switch objType.Kind() {
	case reflect.Ptr:
		// 如果是指针类型，则尝试获取其指向的内容
		if objVal.CanInterface() && !objVal.IsNil() {
			return ToString(objVal.Elem().Interface())
		}
	case reflect.Struct:
		// 针对 time.Time 类型做特殊处理
		if objVal.CanInterface() {
			if _, ok := objVal.Interface().(time.Time); ok {
				break
			}
		}

		buffer := bytes.NewBufferString("{ ")
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			fVal := objVal.Field(i)

			// 判断是否为私有成员
			if IsPrivateMember(field.Name) {
				buffer.WriteString(fmt.Sprintf("'%v':<Private> ", field.Name))
				continue
			}

			switch field.Type.Kind() {
			case reflect.Struct:
				if fVal.CanInterface() {
					// 针对 time.Time 类型做特殊处理
					if t, ok := fVal.Interface().(time.Time); ok {
						buffer.WriteString(fmt.Sprintf("'%v':%v ", field.Name, t))
						break
					}
				}

				// 如果该字段为Struct，则递归调用ToString方法
				buffer.WriteString(fmt.Sprintf("'%v':%v ", field.Name, ToString((fVal.Interface()))))
			case reflect.Slice, reflect.Array, reflect.Map:
				buffer.WriteString(fmt.Sprintf("'%v':%v ", field.Name, ToString((fVal.Interface()))))
			case reflect.Ptr:
				if fVal.CanInterface() {
					var val string
					if fVal.IsNil() {
						val = "<Nil>"
					} else if !fVal.Elem().IsValid() {
						val = "<Invalid>"
					} else {
						val = ToString(fVal.Elem().Interface())
					}
					buffer.WriteString(fmt.Sprintf("'%v':%v ", field.Name, val))
				}
			default:
				buffer.WriteString(fmt.Sprintf("'%v':%v ", field.Name, fVal.Interface()))
			}
		}
		buffer.WriteString("}")
		return buffer.String()
	case reflect.Slice, reflect.Array:
		buffer := bytes.NewBufferString("[ ")
		if objVal.CanInterface() && objVal.IsValid() {
			// 遍历数组元素
			for i := 0; i < objVal.Len(); i++ {
				switch objVal.Index(i).Kind() {
				case reflect.Struct, reflect.Array, reflect.Slice:
					buffer.WriteString(fmt.Sprintf("%v ", ToString(objVal.Index(i).Interface())))
				case reflect.Ptr:
					if objVal.Index(i).CanInterface() {
						buffer.WriteString(fmt.Sprintf("%v ", ToString(objVal.Index(i).Elem().Interface())))
					}
				default:
					buffer.WriteString(fmt.Sprintf("%v ", objVal.Index(i)))
				}
			}
		}
		buffer.WriteString("] ")
		return buffer.String()
	case reflect.Map:
		buffer := bytes.NewBufferString("{ ")
		if objVal.CanInterface() && objVal.IsValid() {
			for _, key := range objVal.MapKeys() {
				val := objVal.MapIndex(key)
				buffer.WriteString(fmt.Sprintf("'%v':%v ", key, ToString(val.Interface())))
			}
		}
		buffer.WriteString("} ")
		return buffer.String()
	}
	return fmt.Sprintf("%v", obj)
}

// 根据字段名或函数名判断其是否为私有成员
func IsPrivateMember(name string) bool {
	fs := []rune(name)
	if len(fs) > 0 {
		return unicode.IsLower(fs[0])
	}
	return false
}

// 判断是否为空，对于String类型，nil或是空字符串均认为是空。对于Array/Map/Chan/Slice nil或者长度为0均认为是空。
func IsEmpty(obj interface{}) bool {
	// 如果参数为nil，则返回true
	if obj == nil {
		return true
	}

	// 根据传入参数的类型进行判断
	objType := reflect.TypeOf(obj)
	switch objType.Kind() {
	case reflect.Ptr:
		// 如果是指针类型，则尝试判断其指向的内容
		return IsEmpty(reflect.ValueOf(obj).Elem())
	case reflect.Array, reflect.Map, reflect.Chan, reflect.Slice, reflect.String:
		if reflect.ValueOf(obj).Len() == 0 {
			return true
		}
		// TODO
	}
	return false
}
