package repository

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type LogsRabbitMQ struct {
	rmq *amqp.Connection
}

func NewLogsRabbitMQ(rmq *amqp.Connection) *LogsRabbitMQ {
	return &LogsRabbitMQ{rmq: rmq}
}

func (r *LogsRabbitMQ) PublishLog(exchangeName, logType, msg string) error {
    ch, err := r.rmq.Channel()
    if err != nil {
        return err
    }
    defer ch.Close()

    err = ch.ExchangeDeclare(
        exchangeName, // name
        "direct",     // type
        true,         // durable
        false,        // auto-deleted
        false,        // internal
        false,        // no-wait
        nil,          // arguments
    )
    if err != nil {
        return err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err = ch.PublishWithContext(ctx,
        exchangeName,
        logType,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(logType+": "+msg),
        })
    if err != nil {
        return err
    }
    return nil
}

