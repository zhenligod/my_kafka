package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/robfig/config"
	"github.com/zhenligod/customer/customer"
)

var (
	configFile = flag.String("configfile", "conf/kafka.conf", "General configuration file")
	conf       = make(map[string]string)
)

func init() {
	cpuNum := runtime.NumCPU()
	log.Println("current cpu nums: ", cpuNum)
	cfg, err := config.ReadDefault(*configFile) //读取配置文件，并返回其Config

	if err != nil {
		log.Fatalf("Fail to find %v,%v", *configFile, err)
	}
	if cfg.HasSection("kafka_log") { //判断配置文件中是否有section（一级标签）
		options, err := cfg.SectionOptions("kafka_log") //获取一级标签的所有子标签options（只有标签没有值）
		if err == nil {
			for _, v := range options {
				optionValue, err := cfg.String("kafka_log", v) //根据一级标签section和option获取对应的值
				if err == nil {
					conf[v] = optionValue
				}
			}
		}
	}
	log.Println(conf)

}

func main() {
	KafkaCustomerConf := customer.KafkaConf{}
	KafkaCustomerConf.IP = conf["kafka_hostname"]
	KafkaCustomerConf.Port = conf["kafka_port"]
	KafkaCustomerConf.Topic = conf["kafka_topic"]
	customer.Customer(KafkaCustomerConf)
}
