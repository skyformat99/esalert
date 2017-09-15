package esalert

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"encoding/base64"
	"fmt"
)

var esurl, basicAuth string

// intiConfig 根据配置文件路径加载配置
func IntiConfig(config_dir string) (*Config, error) {
	bytes,err := ioutil.ReadFile(config_dir)
	if err != nil {
		return nil, err
	}
	config := &Config{
		Host:"localhsot",
		Port:"9200",
		Username:"elastic",
		Password:"changeme",
		Rules: []Rule{},
	}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		 return nil, err
	}
	return config, nil
}

func Run(config *Config) error {
	if len(config.Rules) == 0 {
		return fmt.Errorf("监控规则不能为空!")
	}
	esurl = config.Host + ":" + config.Port
	basicAuth = base64.StdEncoding.EncodeToString([]byte(config.Username + ":" + config.Password))

	return nil
}