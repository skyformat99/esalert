package esalert

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Rules    []RuleConfig
}

type RuleConfig struct {
	Index    string
	Query    interface{}
	Hits     int
	Interval int
	Alert    AlertConfig
}

type AlertConfig struct {
	Type string
	Url  string
}
