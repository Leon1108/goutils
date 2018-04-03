package apiutils

import "fmt"

type CommonStatus struct {
	Code    CommonStatusCode `json:"code"`
	Message string           `json:"message"`
}

func (cs *CommonStatus) String() string {
	return fmt.Sprintf("Status { Code: %d; Message: %v }", cs.Code, cs.Message)
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

func NewCommonResp(payload interface{}, err error) *CommonResp {
	if err != nil {
		return NewFailedCommonResp(err)
	} else {
		return NewSuccessCommonResp(payload)
	}
}

func NewSuccessCommonResp(payload interface{}) *CommonResp {
	return &CommonResp{&CommonStatus{CommonStatusCodeSuccess, "Success"}, payload}
}

func NewFailedCommonResp(err error) *CommonResp {
	return &CommonResp{&CommonStatus{CommonStatusCodeFailed, err.Error()}, nil}
}
