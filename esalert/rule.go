package esalert

import (
	"log"
	"time"
)

type rule interface {
	run()
}

type sampleRule struct {
	esRequest EsRequest
	tick      *time.Ticker
	hits      int
	alerter   []Alerter
}

func (rule sampleRule) run() {
	go func() {
		for {
			select {
			case <-rule.tick.C:
				hits, err := rule.esRequest.RunQuery()
				if err != nil {
					log.Println(err)
				}
				if hits.Total >= rule.hits {
					for _, alerter := range rule.alerter {
						alerter.alert(hits)
					}
				}
			}
		}
	}()
}
