package akafka

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/kafka_topic_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/service/checkout_service"
)

func OrderProducer(order checkout_service.OrderQueueProducerDTO) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Envs.Kafka.Host,
	})
	if err != nil {
		log.Fatal("error on start kafka order producer")
	}

	defer p.Close()

	topic := kafka_topic_enum.ORDER

	value, err := json.Marshal(order)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, nil)

	if err != nil {
		panic(err)
	}

	p.Flush(1000)
}
