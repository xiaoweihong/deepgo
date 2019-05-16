package redispool

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	URL_QUEUE        = "url_queue"
	URL_ADD_TO_REDIS = "url_add_to_redis"
)

func NewPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 1 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

func AddToSet(url string, p *redis.Pool) {
	_, err := p.Get().Do("SADD", URL_ADD_TO_REDIS, url)
	if err != nil {
		fmt.Println("err")
	}
}

func GetQueueLength(p *redis.Pool) int64 {
	result, err := redis.Int64(p.Get().Do("SCARD", URL_ADD_TO_REDIS))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return result
}

func PopfromQueue(p *redis.Pool) string {
	result, err := redis.String(p.Get().Do("SPOP", URL_ADD_TO_REDIS))
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	p := NewPool("192.168.2.189:6379")
	//AddToSet("232",p)
	AddToSet("44", p)
	fmt.Println()
}
