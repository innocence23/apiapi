package queue

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
)

// nsq发布消息
func InitProducer(msgBody string) {
	// 默认配置
	var Host = viper.GetString("nsq.addr")
	var Topic = viper.GetString("nsq.topic")
	// 新建生产者
	p, err := nsq.NewProducer(Host, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	// 发布消息
	if err := p.Publish(Topic, []byte(msgBody)); err != nil {
		panic(err)
	}
}

// nsq订阅消息
type ConsumerT struct{}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

func InitConsumer() {
	var Host = viper.GetString("nsq.addr")
	var Topic = viper.GetString("nsq.topic")
	var Channel = viper.GetString("nsq.channel")
	// 新建一个消费者
	cfg := nsq.NewConfig()
	//设置重连时间
	cfg.LookupdPollInterval = time.Second
	c, err := nsq.NewConsumer(Topic, Channel, cfg)

	if err != nil {
		panic(err)
	}
	// 添加消息处理
	c.AddHandler(&ConsumerT{})
	// 建立连接
	if err := c.ConnectToNSQD(Host); err != nil {
		panic(err)
	}
	/*
		//建立多个nsqd连接
		if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
			panic(err)
		}
		//建立一个nsqd连接
		if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
			panic(err)
		}
	*/
}

func TestNsq() {
	InitConsumer()
	for i := 0; i < 10; i++ {
		InitProducer("hello " + strconv.Itoa(i))
	}
}
