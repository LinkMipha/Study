package main

import (
	"sync"
	"github.com/Shopify/sarama"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect error:", err)
		return
	}
	fmt.Println("connnect success...")
	defer consumer.Close()
	partitions, err := consumer.Partitions("revolution")
	if err != nil {
		fmt.Println("geet partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("revolution", p, sarama.OffsetOldest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func(){
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}