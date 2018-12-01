package main

import (
	"context"
	"fmt"
	"github.com/wenfengwang/iMQ/baton/pb"
	"github.com/wenfengwang/iMQ/client"
	"time"
)

type TestErr struct {
	err  string
	code int
}

func (err TestErr) Error() string {
	return ""
}

func main() {
	cli := client.NewBatonClient()
	for i := 0; i < 8; i++ {
		go func() {
			for {
				_, err := cli.CreateTopic(context.Background(), &pb.CreateTopicRequest{Name: "test", QueueNumbers: 4})
				//err.(TestErr).code
				switch err.(type) {
				case TestErr:

				}
				if err != nil {
					fmt.Println("error:", err)
				}
				//fmt.Println("response code:", res.ResponseCode)
			}
		}()
	}
	time.Sleep(time.Hour)

}
