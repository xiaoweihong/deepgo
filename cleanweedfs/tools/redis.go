package tools

import (
	"fmt"
	"github.com/monnand/goredis"
)

var (
	client      goredis.Client
	redisClient *goredis.Client
)

const (
	URL_QUEUE        = "url_queue"
	URL_ADD_TO_REDIS = "url_add_to_redis"
)

func Ping() (s string, e error) {
	return client.Ping()
}

func ConnectRedis(addr string) {
	client.Addr = addr

}

func PutQueue(url string) {
	err := client.Lpush(URL_QUEUE, []byte(url))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func PopfromQueue() string {
	res, err := client.Spop(URL_ADD_TO_REDIS)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func AddToSet(url string) {
	client.Sadd(URL_ADD_TO_REDIS, []byte(url))

}

func IsVisit(url string) bool {
	bIsVisit, err := client.Sismember(URL_ADD_TO_REDIS, []byte(url))
	if err != nil {
		return false
	}
	return bIsVisit
}

func GetQueueLength() int {
	length, err := client.Scard(URL_ADD_TO_REDIS)
	if err != nil {
		return 0
	}
	return length
}
