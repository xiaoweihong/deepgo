package tools2

import (
	"fmt"
	. "github.com/go-redis/redis"
)

const (
	URL_QUEUE        = "url_queue"
	URL_ADD_TO_REDIS = "url_add_to_redis"
)

func GetRedisClientPool(addr string) *Client {
	redisdb := NewClient(&Options{
		Addr:     addr,
		Password: "",
		DB:       0,
		PoolSize: 100})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}

func AddToSet(url string, client *Client) {
	client.SAdd(URL_ADD_TO_REDIS, url)

}

func GetQueueLength(addr string, client *Client) int64 {
	card := client.SCard(URL_ADD_TO_REDIS)
	return card.Val()
}

func PopfromQueue(client *Client) string {
	pop := client.SPop(URL_ADD_TO_REDIS)
	return pop.Val()
}
