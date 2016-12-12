package logger

import (
	"encoding/json"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	if conf, err := loadConfigFile("./logger_test.conf"); err != nil {
		t.Fatalf("Error: %v", err)
	} else {
		conf := conf[0].Engine.Conf
		t.Logf(">>> %v", conf)
		if b, err := json.Marshal(conf); err != nil {
			t.Fatalf("Error: %v", err)
		} else {
			t.Logf(">>> %v", string(b))
		}
	}
}
