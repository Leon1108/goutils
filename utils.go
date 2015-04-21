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
)

// 获取当前的毫秒值
func GetCurrentMillisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
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
	objType := reflect.TypeOf(obj)
	objVal := reflect.ValueOf(obj)

	// 识别输入参数类型
	switch objType.Kind() {
	case reflect.Ptr:
		// 如果是指针类型，则尝试获取其指向的内容
		if objVal.CanInterface() {
			return ToString(objVal.Elem().Interface())
		}
	case reflect.Struct:
		buffer := bytes.NewBufferString("{ ")
		for i := 0; i < objType.NumField(); i++ {
			field := objType.Field(i)
			fVal := objVal.Field(i)
			buffer.WriteString(fmt.Sprintf("%v:%v; ", field.Name, fVal.Interface()))
		}
		buffer.WriteString("}")
		return buffer.String()
	}
	return fmt.Sprintf("%v", obj)
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
