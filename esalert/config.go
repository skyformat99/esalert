package esalert

type Config struct {
	Host string
	Port string
	Usrname string
	Password string
	Rules []Rule
}

type Rule struct {
	Query string
	Script string
	Interval int
	Alert Alert
}

type Alert struct {
	Type string

}