package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	brokers := []string{"localhost:29092"}

	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		return
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("Error closing consumer: %v\n", err)
		}
	}()

	topic := "test_topic"

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Printf("Error getting partitions: %v\n", err)
		return
	}

	for partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("Error creating partition consumer: %v\n", err)
			return
		}
		go func(pc sarama.PartitionConsumer) {
			defer func() {
				if err := pc.Close(); err != nil {
					fmt.Printf("Error closing partition consumer: %v\n", err)
				}
			}()
			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("Received message from partition %d: %s\n", msg.Partition, string(msg.Value))
				case err := <-pc.Errors():
					fmt.Printf("Error: %v\n", err)
				}
			}
		}(partitionConsumer)
	}
	select {}

}
