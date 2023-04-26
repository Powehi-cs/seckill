package rabbitMQ

import "github.com/gin-gonic/gin"

func Produce(ctx *gin.Context, msg string) {
	mq := GetRabbitMQ()
	mq.PublishSimple(msg)
}
