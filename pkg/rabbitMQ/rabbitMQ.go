package rabbitMQ

import (
	"fmt"
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	MqUrl string
	// 验证
	confirms chan amqp.Confirmation
}

var mq *RabbitMQ

func GetRabbitMQ() *RabbitMQ {
	return mq
}

func NewMQ(queueName string) {
	mq = NewRabbitMQSimple(queueName)
}

func getURL() string {
	user := viper.GetString("rabbitMQ.user")
	password := viper.GetString("rabbitMQ.password")
	ip := viper.GetString("rabbitMQ.ip")
	port := viper.GetString("rabbitMQ.port")
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, ip, port)
}

// NewRabbitMQ 创建结构体实例
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: getURL()}
}

// Destroy 断开channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

// NewRabbitMQSimple 创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	errors.PrintInStdout(err)
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	errors.PrintInStdout(err)

	rabbitmq.confirms = rabbitmq.channel.NotifyPublish(make(chan amqp.Confirmation))
	err = rabbitmq.channel.Confirm(false)
	errors.PrintInStdout(err)

	return rabbitmq
}

// PublishSimple 直接模式队列生产
func (r *RabbitMQ) PublishSimple(ctx *gin.Context, message string) {
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	errors.PrintInStdout(err)

	//调用channel 发送消息到队列中
	err = r.channel.PublishWithContext(
		ctx,
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	// 等待
	confirmed := <-r.confirms
	if confirmed.Ack {
		log.Println("ack success", confirmed.DeliveryTag)
	} else {
		log.Fatalln("ack error", confirmed.DeliveryTag)
	}
	errors.PrintInStdout(err)
}

// ConsumeSimple simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple() {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		true,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	errors.PrintInStdout(err)
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		false, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	errors.PrintInStdout(err)

	forever := make(chan os.Signal, 2)
	signal.Notify(forever, os.Interrupt, os.Kill)

	//启用协程处理消息
	go func() {
		for d := range msgs {
			select {
			case <-forever:
				log.Fatalln("消费者退出！")
			default:
				log.Printf("Received a message: %s", d.Body)
				err = d.Ack(true)
				errors.PrintInStdout(err)
			}
			log.Println("等待下一个")
		}
	}()
	<-forever
	r.Destroy()
	log.Println("消费者退出")
}
