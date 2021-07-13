package main

import (
	"context"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/rlog"
)

func main() {

	c, err := rocketmq.NewPullConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
	)
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new pullConsumer: %s", err), nil)
	}
	err = c.Start()
	if err != nil {
		rlog.Fatal(fmt.Sprintf("fail to new pullConsumer: %s", err), nil)
	}

	ctx := context.Background()
	queue := primitive.MessageQueue{
		Topic:      "TopicTest",
		BrokerName: "", // replace with your broker name. otherwise, pull will failed.
		QueueId:    0,
	}

	offset := int64(0)
	for {
		resp, err := c.PullFrom(ctx, queue, offset, 10)
		if err != nil {
			if err == rocketmq.ErrRequestTimeout {
				fmt.Printf("timeout \n")
				time.Sleep(1 * time.Second)
				continue
			}
			fmt.Printf("unexpectable err: %v \n", err)
			return
		}
		if resp.Status == primitive.PullFound {
			fmt.Printf("pull message success. nextOffset: %d \n", resp.NextBeginOffset)
			for _, msg := range resp.GetMessageExts() {
				fmt.Printf("pull msg: %v \n", msg)
			}
		}
		offset = resp.NextBeginOffset
	}
}