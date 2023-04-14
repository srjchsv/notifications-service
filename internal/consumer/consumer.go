package consumer

import (
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/srjchsv/notifications-service/internal/notifications"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(kafkaHost string) (*KafkaConsumer, error) {
	kc := &KafkaConsumer{}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    kafkaHost,
		"auto.offset.reset":    "latest",
		"group.id":             1,
		"max.poll.interval.ms": 86400000,
	})

	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
		return kc, err
	}

	kc.consumer = c

	return kc, nil
}

func (kc *KafkaConsumer) ConsumeNotifications(notifService *notifications.NotificationService) {
	topic := "notifications"
	err := kc.consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %v", topic, err)
	}

	for {
		msg, err := kc.consumer.ReadMessage(-1)
		if err == nil {
			var notif notifications.Notification
			if err := json.Unmarshal(msg.Value, &notif); err != nil {
				log.Printf("Failed to unmarshal notification: %v", err)
			} else {
				notifService.Add(&notif)
				log.Printf("Received notification: %+v", notif)
			}
		} else {
			log.Printf("Failed to read message: %v", err)
		}
	}
}
