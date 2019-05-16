package main

import (
	"database/sql"
	"deepgo/cleanweedfs/tools2"
	"flag"
	"fmt"
	. "github.com/go-redis/redis"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/micro/go-config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Image struct {
	ImageUrl    string
	CutboardUri string
}

type Redis struct {
	Addr string `json:"addr"`
}

type Postgres struct {
	Addr     string
	Port     string
	DbUser   string
	DbPasswd string
	DataBase string
}

type TimeRange struct {
	StartTimeTs int64
	EndTimeTs   int64
}

var (
	postgres    Postgres
	redis       Redis
	timeRange   TimeRange
	connStr     string
	closing     = make(chan string)
	db          *sql.DB
	redis_str   string
	startTimeTs int64
	endTimeTs   int64
	count       int64
	wg          sync.WaitGroup
	m           sync.Mutex
	clean_url   string
	pool        *Client
)

func init() {
	flag_init()

	err := config.LoadFile("./config_weedfs.json")
	if err != nil {
		panic("load config file error...")
	}

	err = config.Get("database", "postgres").Scan(&postgres)
	if err != nil {
		panic("get postgres config error...")
	}

	err = config.Get("timerange", "range").Scan(&timeRange)
	if err != nil {
		panic("get timerange config error...")
	}
	startTimeTs = timeRange.StartTimeTs
	endTimeTs = timeRange.EndTimeTs

	fmt.Println(timeRange)
	connStr = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgres.DbUser, postgres.DbPasswd, postgres.Addr, postgres.Port, postgres.DataBase)
	db, _ = sql.Open("postgres", connStr)

	err = config.Get("cache", "redis").Scan(&redis)
	if err != nil {
		panic("get redis config error...")
	}
	redis_str = redis.Addr
	pool = tools2.GetRedisClientPool(redis_str)
	clean_url = fmt.Sprintf("http://%s:9333/vol/vacuum", postgres.Addr)
	//使用Ping检查数据库是否实际可用
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func flag_init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("logtostderr", "true")
	flag.Set("v", "1")
}

func converToUrl(url string) (resultUrl string) {
	resultUrl = strings.Replace(url, "8501/api/file", "9333", -1)
	return
}

func httpDo(url string, methodType string, param []byte) (result string, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(methodType, url, strings.NewReader(string(param)))
	if err != nil {
		return
	}
	response, err := client.Do(request)

	if err != nil {

		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("--------->", err)
		return
	}
	result = string(body)

	return

}

func saveImageByType(objectType string, startTimeTs int64, endTimeTs int64) {
	image := Image{}
	rows, err := db.Query("select image_uri,cutboard_image_uri from "+objectType+" where ts > $1 AND ts <= $2", startTimeTs, endTimeTs)
	glog.V(1).Infof("select image_uri,cutboard_image_uri from "+objectType+" where ts > %d AND ts <= %d \n", startTimeTs, endTimeTs)
	defer rows.Close()
	defer wg.Done()
	if err != nil {
		log.Fatal("www", err)
	}

	for rows.Next() {
		err := rows.Scan(&image.ImageUrl, &image.CutboardUri)
		if err != nil {
			panic(err)
		}
		imageUrl := converToUrl(image.ImageUrl)
		cutBoardUrl := converToUrl(image.CutboardUri)
		glog.V(2).Infof("%v full     image--->%v", objectType, imageUrl)
		glog.V(2).Infof("%v cutboard image--->%v", objectType, cutBoardUrl)
		tools2.AddToSet(imageUrl, pool)
		tools2.AddToSet(cutBoardUrl, pool)

		m.Lock()
		count += 1
		m.Unlock()

		glog.V(1).Infof("save %d th %v to redis ", count, objectType)
	}

}

func deleteImageByType(startTimeTs int64, endTimeTs int64) {
	wg.Add(4)

	go deleteRecordByType("vehicles", startTimeTs, endTimeTs)
	go deleteRecordByType("faces", startTimeTs, endTimeTs)
	go deleteRecordByType("nonmotors", startTimeTs, endTimeTs)
	go deleteRecordByType("pedestrians", startTimeTs, endTimeTs)

}

func deleteRecordByType(objectType string, startTimeTs int64, endTimeTs int64) {
	stmt, err := db.Prepare("delete from " + objectType + " where ts > $1 and ts < $2")
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(startTimeTs, endTimeTs)
	resultNum, _ := result.RowsAffected()
	defer db.Close()
	if err != nil {
		panic(err)
	} else if resultNum == 0 {
		glog.V(1).Infoln("record is empty")
	} else {
		glog.V(1).Infoln("delete successful")
	}
	defer wg.Done()
}

func deleteImage() {

	for {
		url := tools2.PopfromQueue(pool)
		if url == "" {
			glog.V(1).Infoln("url delete done")
			break
		}
		resp, err := http.Head(url)
		if err != nil {

		}
		resultUrl := resp.Request.URL.String()
		glog.V(1).Infoln("delete URL", resultUrl)
		//httpDo(resultUrl, "DELETE", []byte(""))
		wg.Add(1)
		go httpDo_batch(resultUrl, "DELETE", []byte(""), &wg)
	}
}

func httpDo_batch(url string, methodType string, param []byte, wg *sync.WaitGroup) (result string, err error) {

	client := &http.Client{}
	request, err := http.NewRequest(methodType, url, strings.NewReader(string(param)))
	if err != nil {
		return
	}
	response, err := client.Do(request)

	if err != nil {

		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("--------->", err)
		return
	}
	result = string(body)
	defer wg.Done()
	return
}

func main() {
	a := time.Now()
	flag.Parse()
	var object_list = []string{"vehicles", "faces", "nonmotors", "pedestrians"}

	wg.Add(4)
	for _, object := range object_list {
		go saveImageByType(object, startTimeTs, endTimeTs)
	}
	wg.Wait()

	deleteImage()
	deleteImageByType(startTimeTs, endTimeTs)
	_, err := httpDo(clean_url, "GET", []byte(""))
	if err != nil {
		glog.Error("err=", err)
	} else {
		glog.V(1).Infoln("clean get successful")
	}
	glog.V(1).Infoln("cost--->", time.Since(a))
	glog.V(1).Infoln("total--->", count)
}
