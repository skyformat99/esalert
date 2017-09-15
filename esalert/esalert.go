package esalert

import (
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var esurl, basicAuth string

// IntiConfig 根据配置文件路径加载配置
func IntiConfig(configDir string) (*Config, error) {
	bytes, err := ioutil.ReadFile(configDir)
	if err != nil {
		return nil, err
	}
	config := &Config{
		Host:     "localhsot",
		Port:     "9200",
		Username: "elastic",
		Password: "changeme",
		Rules:    []RuleConfig{},
	}
	err = yaml.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Run 启动配置参数
func Run(config *Config) error {
	if len(config.Rules) == 0 {
		return fmt.Errorf("监控规则不能为空!")
	}
	rules := []rule{}
	for _, rule := range config.Rules {
		rules = append(rules, sampleRule{
			esRequest: getEsRequest(*config, rule),
			tick:      time.NewTicker(time.Duration(rule.Interval) * time.Second),
			hits:      rule.Hits,
			alerter:   getAlert(rule.Alert),
		})
	}
	for _, rule := range rules {
		rule.run()
	}
	return nil
}

func getAlert(alert AlertConfig) Alerter {
	return HttpAlert{
		url: alert.Url,
	}
}

func getEsRequest(config Config, rule RuleConfig) EsRequest {
	return EsRequest{
		host:     config.Host,
		port:     config.Port,
		name:     config.Username,
		password: config.Password,
		index:    rule.Index,
		query:    rule.Query,
	}
}