package apiutils

import (
	"reflect"
	"testing"
)

func TestNewSuccessCommonResp(t *testing.T) {
	type args struct {
		payload interface{}
	}
	tests := []struct {
		name string
		args args
		want *CommonResp
	}{
		{
			name: "test01",
			args: args{
				payload: &ListPayload{},
			},
			want: &CommonResp{Status: &CommonStatus{PayloadType: PayloadTypeList}},
		},
		{
			name: "test02",
			args: args{
				payload: ListPayload{},
			},
			want: &CommonResp{Status: &CommonStatus{PayloadType: PayloadTypeList}},
		},
		{
			name: "test03",
			args: args{
				payload: nil,
			},
			want: &CommonResp{Status: &CommonStatus{PayloadType: PayloadTypeNormal}},
		},
		{
			name: "test04",
			args: args{
				payload: map[string]interface{}{"name": "value"},
			},
			want: &CommonResp{Status: &CommonStatus{PayloadType: PayloadTypeNormal}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSuccessCommonResp(tt.args.payload); !reflect.DeepEqual(got.Status.PayloadType, tt.want.Status.PayloadType) {
				t.Errorf("NewSuccessCommonResp() = %v, want %v", got, tt.want)
			}
		})
	}
}
