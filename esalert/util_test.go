package esalert

import "testing"

var (
	TestMail = Mail{
		Host:     "smtp.exmail.qq.com",
		Port:     "25",
		Username: "***",
		Password: "fudali133B",
		ReplyTo:  "fuyi@23mofang.com",
	}
	SendTo = []string{"fuyi@23mofang.com"}
)

func TestMail_Send(t *testing.T) {
	/*mail := TestMail
	err := mail.Send(SendTo, "test", []byte("test llll 啪啪啪啪啪  \r\n  sfsdfdfdss saffasafsafafsbdhfhfd"))
	if err != nil {
		t.Error(err)
	}
	return*/
}
