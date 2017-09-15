package main

import (
	"esalert"
	"flag"
	"log"
	"os"
)

var config string

func main() {
	flag.StringVar(&config, "config", "config.yml", "配置文件")
	config, err := esalert.IntiConfig(config)
	checkErr(err)
	err = esalert.Run(config)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Print("解析配置文件出错, ", err)
		os.Exit(1)
	}
}
