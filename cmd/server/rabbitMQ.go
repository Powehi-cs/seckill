package main

import (
	"fmt"
	"github.com/Powehi-cs/seckill/internal/config"
	"github.com/Powehi-cs/seckill/pkg/database"
	"github.com/Powehi-cs/seckill/pkg/errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

var queueName string

func init() {
	config.ReadConfig()
	database.MySQLConnect()
}

func main() {
	url := getURL()
	conn, err := amqp.Dial(url)
	errors.PrintInStdout(err)

	defer conn.Close()

	ch, err := conn.Channel()
	errors.PrintInStdout(err)

	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	errors.PrintInStdout(err)

	err = ch.Qos(
		1,
		0,
		false,
	)
	errors.PrintInStdout(err)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	errors.PrintInStdout(err)

	go func() {
		for msg := range msgs {
			processMSG(msg.Body)
			err = msg.Ack(false)
			errors.PrintInStdout(err)
		}
	}()

}

// 处理消息队列中的消息
func processMSG(msg []byte) {
	//todo: 通过msg更新数据库库存信息
}

func getURL() string {
	rabbitMQ := viper.Get("rabbitMQ").(map[string]string)
	ip := rabbitMQ["ip"]
	port := rabbitMQ["port"]
	user := rabbitMQ["user"]
	password := rabbitMQ["password"]
	queueName = rabbitMQ["name"]
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", ip, port, user, password)
}
