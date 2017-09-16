package esalert

import (
	"log"
	"testing"
)

func Test_IntiConfig(t *testing.T) {
	config, err := IntiConfig("../sample-config.yml")
	if err != nil {
		t.Error(err)
	}
	if len(config.Rules) == 0 {
		t.Error("解析出错")
	}
	for _, rule := range config.Rules {
		if rule.Query == nil {
			t.Error("解析出错")
		}
		json, err := QueryToJSON(rule.Query)
		log.Println(string(json), err)
	}

}
