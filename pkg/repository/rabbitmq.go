package repository

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
)

const(
	ToLogs = "UpdStore.Logs"
	ToErrors = "UpdStore.Errors"
	Info = "info"
	Warning = "warning"
	Error = "error"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewRabbitMQ(cfg RabbitMQConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.Username, cfg.Password, cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

