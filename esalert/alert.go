package esalert

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Alerter 报警方式处理接口
type Alerter interface {
	alert(res Hits) error
}

// LogAlert 打印日志报警方式，默认报警方式，当没有任何报警方式时，自动添加该报警方式
type LogAlert struct {
}

func (LogAlert) alert(hits Hits) error {
	log.Println(hits)
	return nil
}

// HTTPAlert http 报警方式
type HTTPAlert struct {
	url string
}

func (httpAlert HTTPAlert) alert(hits Hits) error {
	buffer := &bytes.Buffer{}
	bytes, _ := json.Marshal(hits)
	buffer.Write(bytes)
	_, err := http.Post(httpAlert.url, "application/josn", buffer)
	if err != nil {
		log.Print("http alert 请求出错,", err)
	}
	return nil
}

// MailAlert 发送邮件报警方式
type MailAlert struct {
	mail MailConfig
}

// TODO
func (mailAlert MailAlert) alert(hits Hits) error {
	SendMail()
	return nil
}
