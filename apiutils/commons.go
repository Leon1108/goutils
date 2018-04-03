package apiutils

import (
	"fmt"
	"reflect"
)

type payloadType string

const (
	PayloadTypeNormal payloadType = "normal" // 普通类型，就是个json object
	PayloadTypeList   payloadType = "list"   // 列表类型，ListPayload
)

type CommonStatus struct {
	Code        CommonStatusCode `json:"code"`
	Message     string           `json:"message"`
	PayloadType payloadType      `json:"payload,omitempty"`
}

func (cs *CommonStatus) String() string {
	return fmt.Sprintf("Status { Code: %d; Message: %v; Payload: %v }", cs.Code, cs.Message, cs.PayloadType)
}

type CommonStatusCode int

const (
	CommonStatusCodeSuccess CommonStatusCode = 0
	CommonStatusCodeFailed  CommonStatusCode = 1
)

type CommonResp struct {
	Status  *CommonStatus `json:"status"`
	Payload interface{}   `json:"payload,omitempty"`
}

type Pagination struct {
	Total    int `json:"total"`  // 总页数
	PageSize int `json:"size"`   // 每页大小, limit
	Offset   int `json:"offset"` // 当前偏移量
}

type ListPayload struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func NewCommonResp(payload interface{}, err error) *CommonResp {
	if err != nil {
		return NewFailedCommonResp(err)
	} else {
		return NewSuccessCommonResp(payload)
	}
}

func NewListResponse(payload *ListPayload) *CommonResp {
	return NewSuccessCommonResp(payload)
}

func NewSuccessCommonResp(payload interface{}) *CommonResp {
	payloadType := PayloadTypeNormal
	if payload != nil {
		payloadTypeName := reflect.TypeOf(payload).Name()
		if reflect.TypeOf(payload).Kind() == reflect.Ptr {
			payloadTypeName = reflect.TypeOf(payload).Elem().Name()
		}

		if payloadTypeName == "ListPayload" {
			payloadType = PayloadTypeList
		}
	}
	return &CommonResp{&CommonStatus{CommonStatusCodeSuccess, "Success", payloadType}, payload}
}

func NewFailedCommonResp(err error) *CommonResp {
	return &CommonResp{&CommonStatus{CommonStatusCodeFailed, err.Error(), ""}, nil}
}
