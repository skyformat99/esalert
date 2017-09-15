package esalert

import (
	"testing"
)

var testErRequest = EsRequest{
	host:     "localhost",
	port:     "9200",
	name:     "elatic",
	password: "changme",
	index:    "logstash-*",
}

func TestEsRequest_getUrl(t *testing.T) {
	url := testErRequest.getUrl()
	if url != "http://localhost:9200/logstash-*/_search" {
		t.Error(url)
	}
}

func TestEsRequest_RunQuery(t *testing.T) {
	hits, err := testErRequest.RunQuery()
	if err != nil {
		t.Error(hits, err)
	}
	if hits.Total < 10 {
		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
}
