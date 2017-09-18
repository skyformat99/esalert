package esalert

import (
	"context"
	"fmt"
	"log"
	"time"
)

type rule interface {
	Name() string
	run(ctx context.Context)
}

type sampleRule struct {
	name      string
	esRequest EsRequest
	tick      *time.Ticker
	time      int32
	hits      int
	alerter   []Alerter
}

func (rule sampleRule) Name() string {
	return rule.name
}

func (rule sampleRule) run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-rule.tick.C:
				log.Println("INFO ", fmt.Sprintf("rule: %s runing", rule.name))
				hits, err := rule.esRequest.RunQuery()
				if err != nil {
					log.Println("ERROR ", err)
					continue
				}
				if hits.Total >= rule.hits {
					for _, alerter := range rule.alerter {
						err := alerter.Alert(hits)
						if err != nil {
							log.Println("ERROR", err)
							// continue
						}
					}
				}
				log.Println("INFO ", fmt.Sprintf("rule: %s run success", rule.name))
			case <-ctx.Done():
				log.Println("INFO ", fmt.Sprintf("rule: %s stoped", rule.name))
				break
			}
		}
	}()
}
