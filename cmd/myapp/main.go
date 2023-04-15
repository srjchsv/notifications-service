package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srjchsv/notifications-service/internal/consumer"
	"github.com/srjchsv/notifications-service/internal/notifications"
	"github.com/srjchsv/notifications-service/internal/pkg/appmetrics"
	"github.com/srjchsv/notifications-service/internal/pkg/ws"
)

func main() {
	kafkaConsumer, err := consumer.NewKafkaConsumer(os.Getenv("KAFKA_HOST"))
	if err != nil {
		log.Fatalf("Failed to initialize Kafka consumer: %v", err)
	}

	notificationsService := notifications.New()
	go kafkaConsumer.ConsumeNotifications(notificationsService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init metrics
	appmetrics.InitPrometheus(e)

	e.GET("/ws", func(c echo.Context) error {
		return ws.UpgradeConnection(c.Response(), c.Request(), notificationsService)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
