////////////////////////////
// Incomplete! Don't use! //
////////////////////////////
package goutils

import (
	"reflect"
	"testing"

	"github.com/Leon1108/goutils_dev/gengo/test"
)

func TestThriftLoggerWriter(t *testing.T) {
	fileName := "xxx.tlog"
	fileDir := "tlogs"
	var uinfo *test.UserInfo
	logger := NewThriftLogger(reflect.TypeOf(uinfo), fileName, fileDir, THRIFT_LOGGER_ROTATE_DAILY)
	t.Logf("%v", logger)

	uinfo = &test.UserInfo{"Leon", "pw1", 18}
	logger.LogNow(uinfo)
}
