package goutils

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

const (
	LEON_STRUCT_TAG_PARAM_NAME = "key"
)

// 根据 tag 名称，在给定的结构体中查找该字段信息，并返回。
// 如果没有找到则会返回一个空的结构（零值？）
// ....
//   UserName	string	`key:"u"`
// ....
func findFiledByTag(tag string, t reflect.Type) (field reflect.StructField, err error) {

	// 验证参数可用性
	if t.Kind().String() != "struct" {
		// TODO 抛出异常
		err = errors.New("The 'Type' is not a struct.")
		return
	}

	// 查找
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get(LEON_STRUCT_TAG_PARAM_NAME) == tag { // Found
			field = t.Field(i)
			return
		}
	}
	// 查找结束，没有找到
	err = errors.New("Not Found!")
	return
}

func ToObject(query string, p2value reflect.Value) interface{} {
	if strings.HasPrefix(query, "/") {
		query = query[1:]
	}

	// 解析查询字符串
	values, err := url.ParseQuery(query)
	if err != nil {
		// TODO 如果解析错误
	}

	// 遍历请求参数
	for k, v := range values {
		field, err := findFiledByTag(k, p2value.Elem().Type())
		if err == nil {
			actualVal := p2value.Elem().FieldByName(field.Name)
			actualVal.SetString(v[0])
		} else {
			// NotFound
		}
	}

	return p2value.Elem().Interface()
}

// 将一个对象转为一个URL查询字符串，并对值进行Escape编码
func ToUrlQueryString(obj interface{}) (str string) {
	val := reflect.ValueOf(obj)

	// 识别传入的对象类型，是否被支持
	switch val.Kind() {
	case reflect.Ptr:
		// 如果传入的是一个指针，则获取其所指向的值
		if "ptr" == val.Kind().String() {
			val = val.Elem()
		}
	case reflect.Struct:
	default:
		return //这里面的都不支持，直接返回空字符串
	}

	// 遍历左右字段
	for i := 0; i < val.NumField(); i++ {
		sf := val.Type().Field(i)                                         // struct field
		if sf.Tag != "" && sf.Tag.Get(LEON_STRUCT_TAG_PARAM_NAME) != "" { // 有 tag，并且tag中包含名为key的key ......
			k := sf.Tag.Get("key")

			str = str + fmt.Sprintf("%v=%v&", k, url.QueryEscape(val.Field(i).String()))
		} else { // 如果没有tag则直接使用字段名
			str = str + fmt.Sprintf("%v=%v&", sf.Name, url.QueryEscape(val.Field(i).String()))
		}
	}
	return str[:len(str)-1] // 截掉最后一个&，并返回
}
