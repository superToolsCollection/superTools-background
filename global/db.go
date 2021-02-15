package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/olivere/elastic/v7"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

/**
* @Author: super
* @Date: 2020-09-18 08:51
* @Description: 全局配置DB
**/
type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

var (
	DBEngine       *gorm.DB
	RedisEngine    *redis.Client
	RabbitMQEngine *RabbitMQ
	ElasticEngine  *elastic.Client
	MongoDBEngine  *mongo.Client
)
