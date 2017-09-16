package esalert

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
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
		return ConfigError{"rules不能为空"}
	}
	rules := []rule{}
	for _, rule := range config.Rules {
		rules = append(rules, sampleRule{
			esRequest: getEsRequest(*config, rule),
			tick:      time.NewTicker(time.Duration(rule.Interval.GetSecond()) * time.Second),
			time:      getTime(rule),
			hits:      rule.Hits,
			alerter:   getAlerts(rule.Alerts),
		})
	}
	for _, rule := range rules {
		rule.run()
	}
	return nil
}

func getTime(rule RuleConfig) int32 {
	res := rule.Time.GetSecond()
	if res == 0 {
		res = rule.Interval.GetSecond()
	}
	return res
}

func getAlerts(alertConfigs []AlertConfig) []Alerter {
	alerterList := make([]Alerter, len(alertConfigs))
	for _, alertConfig := range alertConfigs {
		switch alertConfig.Type {
		case "http":
			alerterList = append(alerterList, HTTPAlert{
				url: alertConfig.URL,
			})
		}
	}
	if len(alerterList) == 0 {
		alerterList = append(alerterList, LogAlert{})
	}
	return alerterList
}

func getEsRequest(config Config, rule RuleConfig) EsRequest {
	return EsRequest{
		host:     config.Host,
		port:     config.Port,
		name:     config.Username,
		password: config.Password,
		index:    rule.Index,
		query:    cleanupMapValue(rule.Query),
	}
}
