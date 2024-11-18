package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {
	server := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "frankhello",
		Value: sarama.StringEncoder("Jang Hello"),
	}

	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("partition: %v offset: %v\n", partition, offset)

}
