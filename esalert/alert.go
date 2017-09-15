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

// HttpAlert http 报警方式
type HttpAlert struct {
	url string
}

func (httpAlert HttpAlert) alert(hits Hits) error {
	buffer := &bytes.Buffer{}
	bytes, _ := json.Marshal(hits)
	buffer.Write(bytes)
	_, err := http.Post(httpAlert.url, "application/josn", buffer)
	if err != nil {
		log.Print("http alert 请求出错,", err)
	}
	return nil
}
