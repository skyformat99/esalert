package esalert

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Rules    []Rule
}

type Rule struct {
	Index    string
	Query    interface{}
	Hits     int
	Interval int
	Alert    Alert
}

type Alert struct {
	Type string
}
