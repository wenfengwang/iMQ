package main

import (
	"github.com/wenfengwang/iMQ/client"
	"flag"
)

var (
	topicName        = flag.String("topic", "testTopic", "topic name")
	batonAddr   = flag.String("baton-addr", "localhost:30000", "address register to baton")
)


func main()  {
	flag.Parse()
	client.BatonAddress = *batonAddr
	//topic := client.NewTopic("testTopic")

}
