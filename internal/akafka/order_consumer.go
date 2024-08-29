package akafka

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/config"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/kafka_group_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/enum/kafka_topic_enum"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/domain/exception"
	"github.com/wesleyfebarretos/ticket-sale/internal/api/logger"
)

func OrderConsumer() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Envs.Kafka.Host,
		"group.id":          kafka_group_enum.ORDER,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		log.Fatalf("error on started kafka order consumer %v", err)
	}

	topic := kafka_topic_enum.ORDER

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(exception.InternalServerException(err.Error()))
	}

	for {
		message, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Println(message.TopicPartition, string(message.Value))
			continue
		}

		if !err.(kafka.Error).IsTimeout() {
			l := logger.Get()

			l.Error().Err(err).Msg("order consumer queue")
		}
	}
}
