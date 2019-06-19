package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = 5
	consumer, err := nsq.NewConsumer("test_topic", "tail", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(func(msg *nsq.Message) error {
		body := append(msg.Body, '\n')
		_, err := os.Stdout.Write(body)
		return err
	}))
	consumer.ConnectToNSQLookupds([]string{"127.0.0.1:4161"})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	consumer.Stop()
	<-consumer.StopChan
}
