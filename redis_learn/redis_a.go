package main

import (
	"deepgo/redis_learn/models"
	"fmt"
	. "github.com/go-redis/redis"
	"sync"
	"time"
)

func main() {
	fmt.Println("-----------------------welcome to redisdemo-----------------------")
	//StringDemo()
	//ListDemo()
	//HashDemo()
	connectPoolTest()
}

func StringDemo() {
	fmt.Println("-----------------------welcome to StringDemo-----------------------")
	redisClient := GetRedisClient()
	if redisClient == nil {
		fmt.Errorf("StringDemo redisClient is nil")
		return
	}

	name := "张三"
	key := "name:zhangsan"
	redisClient.Set(key, name, 1*time.Second)
	val := redisClient.Get(key)
	if val == nil {
		fmt.Errorf("StringDemo get error")
	}
	fmt.Println("name", val)
}

func ListDemo() {
	fmt.Println("-----------------------welcome to ListDemo-----------------------")
	redisClient := GetRedisClient()
	if redisClient == nil {
		fmt.Errorf("ListDemo redisClient is nil")
		return
	}
	articleKey := "article"
	result, err := redisClient.RPush(articleKey, "a", "b", "c").Result() //在名称为 key 的list尾添加一个值为value的元素
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)

	result, err = redisClient.LPush(articleKey, "d").Result() //在名称为 key 的list头添加一个值为value的元素
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:", result)

	length, err := redisClient.LLen(articleKey).Result()
	if err != nil {
		fmt.Println("ListDemo LLen is nil")
	}
	fmt.Println("length: ", length) // 长度

	mapOut, err1 := redisClient.LRange(articleKey, 0, 100).Result()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
}

func HashDemo() {
	fmt.Println("-----------------------welcome to HashDemo-----------------------")
	redisClient := GetRedisClient()
	if redisClient == nil {
		fmt.Errorf("HashDemo redisClient is nil")
		return
	}
	article := models.Article{18, "测试文章内容22222", "测试文章内容22222测试文章内容22222测试文章内容22222", 10, 0}
	articleKey := "article:18"

	redisClient.HMSet(articleKey, models.ToStringDictionary(&article))
	mapOut := redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

	redisClient.HSet(articleKey, "Content", "测试文章内容")
	mapOut = redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

	view, err := redisClient.HIncrBy(articleKey, "Views", 1).Result()
	if err != nil {
		fmt.Printf("\n HIncrBy error=%s ", err)
	} else {
		fmt.Printf("\n HIncrBy Views=%d ", view)
	}
	fmt.Print("\n")

	mapOut = redisClient.HGetAll(articleKey).Val()
	for inx, item := range mapOut {
		fmt.Printf("\n %s:%s", inx, item)
	}
	fmt.Print("\n")

}

func GetRedisClient() *Client {
	redisdb := NewClient(&Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}

func GetRedisClientPool() *Client {
	redisdb := NewClient(&Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
		PoolSize: 5})

	pong, err := redisdb.Ping().Result()
	if err != nil {
		fmt.Println(pong, err)
	}
	return redisdb
}

// 连接池测试
func connectPoolTest() {
	fmt.Println("-----------------------welcome to connect Pool Test-----------------------")
	client := GetRedisClientPool()
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d, IdleConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().IdleConns)
		}()
	}

	wg.Wait()
}
