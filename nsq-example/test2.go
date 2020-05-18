/*
 1 producer and 3 consumer
 produce3() publish "x", "y" to topic "topic_test"
 consume2() subscribe channel "channel_sensor01" of topic "topic_test"
 consume3() subscribe channel "channel_sensor01" of topic "topic_test"
 consume4() subscribe channel "channel_sensor02" of topic "topic_test"
 consume1接收到y，consume2接收到x,z，consume3接收到x,y,z
*/

package example

import (
	"log"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

func TestNSQ2(t *testing.T) {
	NSQDsAddrs := []string{"127.0.0.1:4150"}
	go consume2(NSQDsAddrs)
	go consume3(NSQDsAddrs)
	go consume4(NSQDsAddrs)
	go produce3()
	time.Sleep(5 * time.Second)
}

func produce3() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := producer.Publish("topic_test", []byte("x")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
	if err := producer.Publish("topic_test", []byte("y")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
	if err := producer.Publish("topic_test", []byte("z")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
}

func consume2(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic_test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C1")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C2")
	}
	<-consumer.StopChan
}

func consume3(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic_test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C2")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C3")
	}
	<-consumer.StopChan
}

func consume4(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic_test", "sensor02", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C2")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C4")
	}
	<-consumer.StopChan
}
