package esalert

// Config 全部配置
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Rules    []RuleConfig
	Mail     MailConfig
}

// MailConfig 邮箱信息配置
type MailConfig struct {
	Username string
	Password string
	SMTPHost string   `yaml:"smtp_host"`
	SMTPPort string   `yaml:"smtp_port"`
	SMTPSSL  bool     `yaml:"smtp_ssl"`
	SendTo   []string `yaml:"send_to"`
	FromAddr string   `yaml:"from_addr"`
	ReplyTo  string   `yaml:"reply_to"`
	TPLFile  string   `yaml:"tpl_file"`
	Content  string
}

// RuleConfig 规则配置
type RuleConfig struct {
	Name     string
	Index    string
	Query    interface{}
	Hits     int
	Interval int
	Alerts   []AlertConfig
}

// AlertConfig 报警配置
type AlertConfig struct {
	Type string
	URL  string `yaml:"url"`
	Mail MailConfig
}
