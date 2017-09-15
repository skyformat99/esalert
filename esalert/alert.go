package esalert

import (
	"net/http"
	"encoding/json"
	"bytes"
	"github.com/ngaut/log"
)

type alerter interface {
	alert(res map[string]interface{}) error
}

type HttpAlert struct {
	url string
	request *http.Request
}

func(httpAlert HttpAlert) alert(res map[string]interface{}) error  {
	buffer := &bytes.Buffer{}
	bytes,_ := json.Marshal(res)
	buffer.Write(bytes)
	_,err := http.Post(httpAlert.url, "application/josn", buffer)
	if err != nil {
		log.Error("http alert 请求出路,", err)
	}
	return nil
}