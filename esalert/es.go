package esalert

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// EsRequest ealsticsearch请求对象
type EsRequest struct {
	host     string
	port     string
	name     string
	password string
	index    string
	query    interface{}
	request  *http.Request
}

// RunQuery 发起请求
func (er EsRequest) RunQuery() (Hits, error) {
	var err error
	if er.request == nil {
		var body io.Reader
		if er.query != nil {
			body = ToBuffer(er.query)
		}
		er.request, err = http.NewRequest("get", er.getUrl(), body)
		er.request.SetBasicAuth(er.name, er.password)
		er.request.Header.Set("Content-Type", "Application/json")
	}
	if err != nil {
		return Hits{}, err
	}
	response, err := http.DefaultClient.Do(er.request)
	body, err := ioutil.ReadAll(response.Body)
	if response.StatusCode != 200 {
		return Hits{}, fmt.Errorf(string(body))
	}
	log.Println(string(body))
	if err != nil {
		return Hits{}, err
	}
	var res Res
	json.Unmarshal(body, &res)
	return res.Hits, nil
}

// getUrl 获取查询地址
func (er EsRequest) getUrl() string {
	return fmt.Sprintf("http://%s:%s/%s/_search", er.host, er.port, er.index)
}

type Res struct {
	Hits Hits `json:"hits"`
}

// Hits 返回值的封装
type Hits struct {
	Total    int           `json:"total"`
	MaxScore int           `json:"max_score"`
	Hits     []interface{} `json:"hits"`
}
