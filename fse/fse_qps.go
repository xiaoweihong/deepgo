package main

import (
	"bytes"
	"deepgo/fse/bolt"
	"deepgo/fse/dg_fse"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/golang/glog"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var ADDR = flag.String("addr", "127.0.0.1:8010", "ip:port")

var Command = flag.String("cmd", "write", "write/read")
var FeatureLen = flag.Int("fl", 384, " feature length")
var RepoId = flag.String("rid", "world", " repo name")
var QPS = flag.Int("qps", 1, "qps")
var Thread = flag.Int("t", 8, "thread")
var TimeOUT = flag.Int("timeout", 3, "seconds")
var Record = flag.Bool("r", false, "record feature or use record feature")
var RecordDB = flag.String("db", "ranker", "db name")
var MaxCount = flag.Int("max", 0, "max times")
var LocationNum = flag.Int("locnum", 0, "number of locations")
var DayNum = flag.Int("daynum", 0, "number of days")

var (
	done         chan bool
	instance     *bolt.BoltDB
	once         sync.Once
	RAND         *rand.Rand
	RANDLock     sync.Mutex
	record_keys  []string
	record_index = int32(0)
)
var client = &http.Client{
	Timeout: time.Duration(*TimeOUT) * time.Second,
}

func GetInstance() *bolt.BoltDB {
	if instance == nil {
		once.Do(func() {
			if instance == nil {
				instance = bolt.InitDB(*RecordDB)
				if !instance.IfTableExist([]byte(*RecordDB)) {
					instance.NewTable(*RecordDB)
				}
				if *Record && *Command == "read" {
					record_keys, _ = GetInstance().GetTableKeys(*RecordDB)
					glog.Info("read db key ", len(record_keys))
				}
			}
		})
	}
	return instance
}

func GenF(l int) (string, string) {
	if *Record && *Command == "read" {
		index := atomic.AddInt32(&record_index, 1)
		vb, _ := GetInstance().GetValue([]byte(*RecordDB), []byte(record_keys[int(index)%len(record_keys)]))
		return string(record_keys[int(index)%len(record_keys)]), string(vb)
	}

	fs := make([]float32, l)
	sum := float32(0)
	for i := 0; i < l; i++ {
		RANDLock.Lock()
		fs[i] = (RAND.Float32() - 0.5) * 2
		RANDLock.Unlock()
		sum += fs[i] * fs[i]
	}
	sum = float32(math.Sqrt(float64(sum)))
	for i := 0; i < l; i++ {
		fs[i] /= sum
	}
	bytes := make([]byte, 4*l)
	for i := 0; i < l; i++ {
		bits := math.Float32bits(fs[i])
		binary.LittleEndian.PutUint32(bytes[i*4:], bits)
	}
	id_u, _ := uuid.NewV4()
	id := id_u.String()
	f := base64.StdEncoding.EncodeToString(bytes)
	if *Record && *Command == "write" {
		GetInstance().SetValue([]byte(*RecordDB), []byte(id), []byte(f))
	}
	return id, f
}

func Post(v interface{}, url string) (error, []byte) {
	b, _ := json.Marshal(v)
	var body_rs []byte
	var resp *http.Response
	var err error
	if resp, err = client.Post("http://"+*ADDR+url, "application/json", bytes.NewReader(b)); err != nil {
		switch err := err.(type) {
		case net.Error:
			if err.Timeout() {
				glog.Errorln("net error : post time out ")
			}
		default:
			glog.Errorf("can not connect server %v", err)
		}
		return err, nil
	}
	if resp != nil && resp.Body != nil {
		defer func() {
			io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
		}()
	}

	if resp.StatusCode/100 == 2 {
		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return errors.New("body error " + err.Error()), nil
		} else {
			// {
			// 	"Context": {
			// 		"Status": "200",
			// 		"Message": "OK"
			// 	}
			// }
			type Context struct {
				Status  string
				Message string
			}
			type Resp struct {
				Context Context
			}
			m := Resp{}
			if err := json.Unmarshal(body, &m); err != nil {
				return errors.New("body json error " + err.Error()), nil
			}
			if m.Context.Status != "200" {
				return errors.New("!200 " + m.Context.Status + " " + m.Context.Message), nil
			}
			body_rs = body
		}
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		if body == nil {
			body = make([]byte, 0)
		}
		return errors.New(fmt.Sprintf("post fail - %s(%d)", string(body), resp.StatusCode)), nil
	}

	return nil, body_rs
}
func AddRepo(repo_id string, repo_cap int) error {
	request_tpl := `{
		"Context":{
		},
		"Repo": {
			"Operation": 1,
			"RepoId": "b",
			"Level": 1,
			"FeatureLen": 384,
			"FeatureDataType":3,
			"Capacity": 20000,
			"Params": {	"GPUThreads": "[1]"},
			"IndexType": "IDMap2,Flat",
			"NeedAttribute": true
		}
	}`
	request := &dg_fse.RankRepoOpRequest{}
	json.Unmarshal([]byte(request_tpl), request)
	request.Repo.RepoId = repo_id
	request.Repo.Capacity = int32(repo_cap)
	request.Repo.FeatureLen = int32(*FeatureLen)
	err, _ := Post(request, "/rank/repo")
	return err
}

func DelRepo(repo_id string) error {
	request_tpl := `{
		"Context":{
		},
		"Repo": {
		  "RepoId": "gofse_test",
		  "Operation": 2
		}
	  }`
	request := &dg_fse.RankRepoOpRequest{}
	json.Unmarshal([]byte(request_tpl), request)
	request.Repo.RepoId = repo_id
	err, _ := Post(request, "/rank/repo")
	return err
}

func SetFeatureTimeAndLocation(request *dg_fse.RankFeatureOpRequest, n int) {
	totalFeature := 10000000
	locationNum := 50
	totalDays := 100
	if *MaxCount > 0 {
		totalFeature = *MaxCount
	}
	if *LocationNum > 0 {
		locationNum = *LocationNum
	}
	if *DayNum > 0 {
		totalDays = *DayNum
	}
	featurePerDayPerLoc := totalFeature / locationNum / totalDays
	timeStep := time.Hour * 24 / time.Duration(featurePerDayPerLoc) / time.Millisecond

	loc := n % locationNum
	day := n / (featurePerDayPerLoc * locationNum)
	currentTime := time.Duration(day) * 24 * time.Hour / time.Millisecond
	ms := time.Duration(n%(featurePerDayPerLoc*locationNum)/locationNum) * timeStep
	currentTime += ms

	request.Features.ObjectFeatures[0].Location = strconv.Itoa(int(loc))
	request.Features.ObjectFeatures[0].Time = int64(currentTime)

	if n%100000 == 0 {
		fmt.Printf("%d / %d\n", n, totalFeature)
	}
}

func AddFeature(repo_id string, n int) error {
	request_tpl := `{
		"Features": {
		  "Operation": 1,
		  "RepoId": "b",
		  "ObjectFeatures": [
			{
			  "Feature": "",
			  "Attribute": {
				  "Id": "attr-5",
				  "Attributes": {"sex": "female", "age": "20"}
			  },
			  "Time": 0,
			  "Id": "id-3",
			  "Location": "0"
			}
		  ]
		},
		"Context": {
		  "SessionId": "ss_743"
		}
	  }`
	request := &dg_fse.RankFeatureOpRequest{}
	json.Unmarshal([]byte(request_tpl), request)
	request.Features.RepoId = repo_id
	fid, f := GenF(*FeatureLen)
	request.Features.ObjectFeatures[0].Id = repo_id + "-" + fid
	request.Features.ObjectFeatures[0].Feature = f
	request.Features.ObjectFeatures[0].Attribute.Id = repo_id + "-" + fid
	SetFeatureTimeAndLocation(request, n)
	err, _ := Post(request, "/rank/feature")
	return err
}

func Rank(repo_id string, n int) error {
	request_tpl := `{
		"Params":{
			"RepoId":"abc",
			"Normalization":"true",
			"Locations":"0",
			"StartTime":"0",
			"EndTime":"9999999999999",
			"ShowAttributes":"true",
			"MaxCandidates":"3"
		},
		"ObjectFeature": {
		  "Feature": ""
		},
		"Context": {
		  "SessionId": "test123"
		}
	  }`
	request := &dg_fse.RankFeatureRequest{}
	json.Unmarshal([]byte(request_tpl), request)
	request.Params["RepoId"] = repo_id

	id, f := GenF(*FeatureLen)
	request.ObjectFeature.Feature = f

	err, body := Post(request, "/rank")
	if err == nil {
		resp := &dg_fse.RankFeatureResponse{}
		json.Unmarshal(body, resp)
		if len(resp.Candidates) <= 0 {
			return errors.New("rank Candidates is 0 ")
		}
		if len(resp.Candidates[0].Attributes) <= 0 {
			return errors.New("rank Candidates Attributes is nil ")
		}
		if *Record && *Command == "read" {
			if len(resp.Candidates) < 1 {
				return errors.New(" Candidates len < 1 ")
			} else {
				if resp.Candidates[0].Id != repo_id+"-"+id {
					return errors.New(" Top 1 Candidates  " + id + "!=" + resp.Candidates[0].Id)
				}
			}
		}
	}

	return err
}

func Loop(f func(string, int) error) {
	n := int32(0)
	ct := int32(0)
	ct_time := int64(0)
	max_time := int64(0)
	var lock sync.Mutex
	chans := make(chan bool, *Thread*2)
	for index := 0; index < *Thread; index++ {
		go func() {
			for _ = range chans {
				st := time.Now()
				err := f(*RepoId, int(atomic.AddInt32(&n, 1)))
				if err != nil {
					glog.Errorln("do failed ", err.Error())
				} else {
					ct = atomic.AddInt32(&ct, 1)
					ed := time.Now()
					tp := ed.Sub(st).Nanoseconds()
					ct_time = atomic.AddInt64(&ct_time, tp)
					if tp >= max_time {
						lock.Lock()
						if tp >= max_time {
							max_time = tp
						}
						lock.Unlock()
					}
				}
			}
		}()
	}
	tr := time.NewTicker(time.Second / time.Duration(*QPS))
	st := time.Now()
	tr_report := time.NewTicker(time.Second * 10)
	fastCount := 0
	fastCount_last := 0
	maxCount := 0
	for {
		if *MaxCount > 0 && maxCount >= *MaxCount {
			glog.Infoln("count ", maxCount, *MaxCount, " close submit")
			return
		}
		select {
		case <-done:
			close(chans)
			return
		case <-tr.C:
			select {
			case chans <- true:
				maxCount++
				break
			default:
				fastCount++
				// log.Println("qps is too fast")
			}
			break
		case <-tr_report.C:
			ed := time.Now()
			glog.Infof("latest 10s  count:%d qps:%f avg:%fms max:%fms", ct, float64(ct)/ed.Sub(st).Seconds(), float64(ct_time)/float64(ct)/1000000, float64(max_time)/1000000)
			if fastCount > fastCount_last {
				glog.Infof("qps is too high,drop %d request", fastCount-fastCount_last)

			}
			fastCount_last = fastCount
			st = time.Now()
			atomic.StoreInt32(&ct, 0)
			atomic.StoreInt64(&ct_time, 0)
			atomic.StoreInt64(&max_time, 0)
		}
	}
}

func WaitForSignal(sources ...os.Signal) os.Signal {
	var s = make(chan os.Signal, 1)
	defer signal.Stop(s) //the second Ctrl+C will force shutdown

	signal.Notify(s, sources...)
	return <-s //blocked
}

func basic() {
	repo_basic := "gofse_test"
	repo_basic_cap := 20000
	if err := AddRepo(repo_basic, repo_basic_cap); err != nil {
		glog.Errorf("add repo error %v", err)
		return
	}
	glog.Info("AddRepo done")
	if err := AddFeature(repo_basic, 0); err != nil {
		glog.Errorf("AddFeature error %v", err)
		return
	}
	glog.Info("AddFeature done")
	if err := Rank(repo_basic, 0); err != nil {
		glog.Errorf("Rank error %v", err)
		return
	}
	glog.Info("Rank done")
	if err := DelRepo(repo_basic); err != nil {
		glog.Errorf("del repo error %v", err)
		return
	}
	glog.Info("DelRepo done")

	glog.Info("all done")
}

func main() {
	//source := rand.NewSource(time.Now().Unix())
	//RAND = rand.New(source)
	done = make(chan bool)
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	glog.Infoln(*MaxCount)
	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 1000
	cmd := &dg_fse.PingRequest{}
	b, _ := json.Marshal(cmd)
	if _, err := http.Post("http://"+*ADDR, "application/json", bytes.NewReader(b)); err != nil {
		glog.Fatalf("can not ping server %v", err)
	}

	switch *Command {
	case "write":
		go Loop(AddFeature)
		break
	case "read":
		go Loop(Rank)
		break
	case "basic":
		basic()
	default:
		glog.Fatal("bad cmd")
	}

	sig := WaitForSignal(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM) //signal exit: Ctrl+C or ...
	close(done)
	glog.Infoln("Gofse", "Exit", "got signal: %v, trigger to stop system", sig)
	return
}
