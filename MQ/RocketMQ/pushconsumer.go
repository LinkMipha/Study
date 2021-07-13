

package RocketMQ

import (
"context"
"fmt"
"github.com/apache/rocketmq-client-go/v2"
"github.com/apache/rocketmq-client-go/v2/consumer"
"github.com/apache/rocketmq-client-go/v2/primitive"
"os"
"time"
)

func main() {
	nameSrv := []string{"10.100.66.23:9876", "10.100.70.230:9876", "10.100.68.78:9876"}
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("upaycenter_fare_activity_consumer"),//设置消费组名称，消费组由框架组创建
		consumer.WithNameServer(nameSrv), //设置name server地址，<=2.0.0版本，不支持域名
		//分别设置框架组提供的accessKey和secretKey
		consumer.WithCredentials(primitive.Credentials{AccessKey: "paycashier_testmq", SecretKey: "FGtWTMmxW0"}),
		//consumer.WithTrace(&primitive.TraceConfig{NamesrvAddrs: nameSrv}), //启用消息轨迹，把该行代码注释掉，则不启用
	)
	// 订阅topic
	err := c.Subscribe("upaycenter_fare_activity_topic", consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		//  消费订阅的消息，在这里增加自己的业务处理，msgs的条数如采用默认配置，没有消息积压的时候一般1-2条，最多32条（与消息体的大小也有关系）
		// 可以通过设置相关参数，增大一次回调消费msgs的条数，设置比较麻烦，一般场景，建议默认配置即可
		for i := range msgs {
			fmt.Printf("subscribe callback: %v \n", msgs[i])
		}

		// 消费成功，返回这个值。如果认为消息处理失败，可以返回     consumer.ConsumeRetryLater，会进行重试（默认消费方式，就是现在示例代码的消费方式）
		// 如果消费失败的话，当前这一批次msgs的消息都会进行重试，所以在消费代码中要采用一些方式避免重复消费
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}