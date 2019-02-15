package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Resp struct {
	Code     int64
	Msg      string
	Redirect string
}

type ResonseData struct {
	TaskId         int64
	UniqueSourceId string
	RenderUri      string
	SensorName     string
	Status         int64
	UniqueTaskId   string
}
type Result struct {
	Resp
	Data []ResonseData
}

type SourceTreeData struct {
	Resp
	Data TreeData
}

type Sensor struct {
	SensorName     string
	Url            string
	Latitude       float64
	Longitude      float64
	Comment        string
	SensorSn       string
	SensorType     int
	UniqueRepoId   string
	UniqueSensorId string
	Sources        []Source
}
type SourceData struct {
	uri      string
	SourceId string
}

type Repo struct {
	UniqueRepoId string
	Name         string
	SensorCount  int64
}

type Source struct {
	Id       int64
	Type     byte
	Uri      string
	Name     string
	Status   byte
	SourceId string
}

type SensorInfo struct {
	AllSize      int64
	ReturnedSize int64
	Sensors      []Sensor
}

type AllSensor struct {
	Resp
	Data SensorInfo
}

type Task struct {
	UniqueSourceId  string
	TypeId          int64
	DetectTypeIds   []int
	AdditionalInfos map[string]string
}

type CommonQuery struct {
	Offset int32
	Limit  int32
}

type SensorSource struct {
	CommonQuery  CommonQuery
	UniqueRepoId string
	SourceTypes  []int
}

type TreeData struct {
	Id           int64
	RepoId       int64
	UniqueRepoId string
	SensorCount  int64
	Name         string
	Repos        []Repo
	Sensors      []Sensor
}

var (
	ip      = flag.String("ip", "192.168.2.22", "deepvideo ip")
	port    = flag.Int("port", 8899, "deepvideo port")
	file    = flag.String("file", "rtsp.txt", "file path about rtsp list")
	command = flag.String("c", "", "command about api.(get_task_list del_all_task add_vehicle_task add_face_task add_kse_task)")
	repoId  = flag.String("repoId", "", "add task need repoId")
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "./tmp")
	flag.Set("logtostderr", "true")
	flag.Set("v", "1")
}

func main() {
	flag.Parse()

	if *ip == "" {
		printUsageErrorAndExit("no -ip specified. should not empty")
	}

	if *port == 0 {
		printUsageErrorAndExit("no -port specified. should not 0")
	}

	switch *command {
	case "get_task_list":
		get_task_list()
	case "add_vehicle_task":
		if *repoId == "" {
			printUsageErrorAndExit("no -repoId specified.add task need repoId,should not empty")
		}
		add_vehicle_task(getSourceByUniqueRepoId(*repoId))
	case "add_face_task":
		if *repoId == "" {
			printUsageErrorAndExit("no -repoId specified.add task need repoId,should not empty")
		}
		add_face_task(getSourceByUniqueRepoId(*repoId))
	case "add_kse_task":
		if *repoId == "" {
			printUsageErrorAndExit("no -repoId specified.add task need repoId,should not empty")
		}
		add_kse_task(getSourceByUniqueRepoId(*repoId))
	case "del_all_task":
		del_all_task()
	case "add_sys_sensor":
		if *repoId == "" {
			printUsageErrorAndExit("no -repoId specified.add task need repoId,should not empty")
		}
		add_sys_sensor(*repoId)
	default:
		getRepoInfo()
	}

	//add_sys_sensor("root")
	//getRepoList()
	//get_task_list()
	//getSourceIdList()
	//add_vehicle_task(getSourceIdList())
	//get_task_list()
	//del_all_task()
	//add_kse_task(getSourceIdList())
	//add_vehicle_task(getSourceByUniqueRepoId("486d83ce-9934-4998-b5ff-29fb5b37e5cb"))
	//getSourceByUniqueRepoId("486d83ce-9934-4998-b5ff-29fb5b37e5cb")

}

func getRepoInfo() {
	fmt.Println(strings.Repeat("******", 10))
	fmt.Println("repoInfo below:")
	url := fmt.Sprintf("http://%s:%d/api/biz/repos/tree?WithSensor=true&WithSource=true&LimitLevel=1&UniqueRepoId=%s&SourceTypes=3&WithTypeReg=&timestamp=1550197932129", *ip, *port, "root")
	result, err := httpDo(url, "GET", []byte(""))
	if err != nil {

	}
	var sourceTreeData SourceTreeData

	json.Unmarshal([]byte(result), &sourceTreeData)
	rootRepo := Repo{
		UniqueRepoId: "root",
		SensorCount:  sourceTreeData.Data.SensorCount,
		Name:         sourceTreeData.Data.Name,
	}
	repoList := sourceTreeData.Data.Repos
	repoList = append(repoList, rootRepo)
	for _, repo := range repoList {
		fmt.Printf("repoId:%s-----sensorCount:%d-----repoName:%s\n", repo.UniqueRepoId, repo.SensorCount, repo.Name)
	}
	fmt.Println(strings.Repeat("******", 10))
}

func getSourceByUniqueRepoId(uniqueRepoId string) (sourceIdList []string) {
	url := fmt.Sprintf("http://%s:%d/api/biz/repos/tree?WithSensor=true&WithSource=true&LimitLevel=1&UniqueRepoId=%s&SourceTypes=3&WithTypeReg=&timestamp=1550197932129", *ip, *port, uniqueRepoId)

	result, err := httpDo(url, "GET", []byte(""))
	if err != nil {

	}
	var sourceTreeData SourceTreeData

	json.Unmarshal([]byte(result), &sourceTreeData)
	sensorList := sourceTreeData.Data.Sensors
	fmt.Println(len(sensorList))
	for _, sensor := range sensorList {
		sourceIdList = append(sourceIdList, sensor.Sources[0].SourceId)
	}

	return
}

func getRepoList() {
	url := fmt.Sprintf("http://%s:%d/api/biz/sys/sensors/list", *ip, *port)
	param := `{"CommonQuery":{"Offset":0,"Limit":10},"UniqueRepoId":"root","SourceTypes":[3,4]}`
	//fmt.Printf("%T",param)
	resultResponse, err := httpDo(url, "POST", []byte(param))
	if err != nil {
		return
	}
	var result AllSensor
	json.Unmarshal([]byte(resultResponse), &result)
	fmt.Println(result)

}

func getSensorsByUniqueRepoId(uniqueRepoId string) (uniqueSendorList []string) {
	commQuery := CommonQuery{
		Offset: 0,
		Limit:  100,
	}
	url := fmt.Sprintf("http://%s:%d/api/biz/sys/sensors/list", *ip, *port)
	sensorSource := SensorSource{
		CommonQuery:  commQuery,
		UniqueRepoId: uniqueRepoId,
		SourceTypes:  []int{3, 4},
	}
	bytes, err := json.Marshal(sensorSource)
	if err != nil {

	}
	//fmt.Println(string(bytes))
	result, err := httpDo(url, "POST", bytes)
	if err != nil {

	}
	var allSensor AllSensor
	json.Unmarshal([]byte(result), &allSensor)
	sensorList := allSensor.Data.Sensors
	for _, sensor := range sensorList {
		uniqueSendorList = append(uniqueSendorList, sensor.UniqueSensorId)
	}
	fmt.Println(uniqueSendorList)
	return

}

//todo 待完善source信息
//func getSourceIdList() (sourceList []string) {
//	url := fmt.Sprintf("http://%s:%d/api/source", *ip, *port)
//	resp, err := http.Get(url)
//	if err != nil {
//
//	}
//	defer resp.Body.Close()
//	res, err := ioutil.ReadAll(resp.Body)
//	var source Source
//	json.Unmarshal(res, &source)
//	for _,v:=range source.Data{
//		sourceList = append(sourceList, v.SourceId)
//	}
//	fmt.Println(len(sourceList))
//	return
//}

func get_task_list() (uniqueTaskIdList []string, err error) {
	url := fmt.Sprintf("http://%s:%d/api/task", *ip, *port)
	resp, err := http.Get(url)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result Result
	json.Unmarshal(res, &result)

	fmt.Println(strings.Repeat("******", 10))
	fmt.Println("taskID", "\t", "SensorName", "\t", "Status", "\t", "rtsp")
	for i := 0; i < len(result.Data); i++ {
		tmp := result.Data[i]
		fmt.Printf("%-d\t\t %-7s\t\t %-d\t\t\t%-s\n", tmp.TaskId, tmp.SensorName, tmp.Status, tmp.RenderUri)
		uniqueTaskIdList = append(uniqueTaskIdList, tmp.UniqueTaskId)
	}
	fmt.Println(strings.Repeat("******", 10))

	return
}

func add_kse_task(sourceList []string) {
	url := fmt.Sprintf("http://%s:%d/api/tasks", *ip, *port)
	var s []Task
	for _, v := range sourceList {

		taskParam := Task{
			UniqueSourceId:  v,
			TypeId:          3,
			DetectTypeIds:   []int{2011, 2012, 2013, 2015},
			AdditionalInfos: map[string]string{"kse": "true"},
		}
		sensors := append(s, taskParam)
		bytes, err := json.Marshal(sensors)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(bytes))
		result, err := httpDo(url, "POST", bytes)
		if err != nil {

		}
		fmt.Println(result)
	}
}

func add_face_task(sourceList []string) {
	url := fmt.Sprintf("http://%s:%d/api/tasks", *ip, *port)
	var s []Task
	for _, v := range sourceList {

		taskParam := Task{
			UniqueSourceId: v,
			TypeId:         3,
			DetectTypeIds:  []int{2011, 2012, 2013, 2015},
		}
		sensors := append(s, taskParam)
		bytes, err := json.Marshal(sensors)
		if err != nil {
			fmt.Println(err)
		}
		result, err := httpDo(url, "POST", bytes)
		if err != nil {

		}
		fmt.Println(result)
	}
}

func add_vehicle_task(sourceList []string) {
	url := fmt.Sprintf("http://%s:%d/api/tasks", *ip, *port)
	var s []Task
	for _, v := range sourceList {

		taskParam := Task{
			UniqueSourceId: v,
			TypeId:         3,
			DetectTypeIds:  []int{2011, 2012, 2013, 2015},
		}
		sensors := append(s, taskParam)
		bytes, err := json.Marshal(sensors)
		if err != nil {
			fmt.Println(err)
		}
		result, err := httpDo(url, "POST", bytes)
		if err != nil {

		}
		fmt.Println(result)
	}
}

func del_all_task() {

	uniqueTaskIdList, err := get_task_list()
	if err != nil {

	}

	for _, uniqueTaskId := range uniqueTaskIdList {

		url := fmt.Sprintf("http://%s:%d/api/tasks?ids=%s", *ip, *port, uniqueTaskId)
		httpDo(url, "DELETE", []byte(""))

	}

}

func import_file(stream_file string) (res []string, err error) {
	if contentBytes, err := ioutil.ReadFile("rtsp.txt"); err == nil {
		result := strings.Replace(string(contentBytes), "\n", "\n", 1)

		splitResult := strings.Split(result, "\n")
		return splitResult, err
	}

	return
}

func add_sys_sensor(repoId string) {
	url := fmt.Sprintf("http://%s:%d/api/biz/sensors", *ip, *port)
	url_list, _ := import_file(*file)
	for _, sensor := range url_list {
		tmp := strings.Split(sensor, " ")
		//linux上有问题，mac上没有，多个[]
		if len(tmp) == 1 {
			continue
		}
		latitude, _ := strconv.ParseFloat(tmp[2], 32)
		longitude, _ := strconv.ParseFloat(tmp[3], 32)
		sensor_param := Sensor{
			SensorName:   tmp[0],
			Url:          tmp[1],
			Latitude:     latitude,
			Longitude:    longitude,
			SensorType:   1,
			UniqueRepoId: repoId,
		}
		var s []Sensor
		data := append(s, sensor_param)

		bytea, err := json.Marshal(data)

		if err != nil {

		}
		result, err := httpDo(url, "POST", bytea)
		if err != nil {

		}

		fmt.Println(result)
	}
}

func httpDo(url string, methodType string, param []byte) (result string, err error) {
	client := &http.Client{}
	request, err := http.NewRequest(methodType, url, strings.NewReader(string(param)))
	if err != nil {
		return
	}
	request.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQzMDIwNTAsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUzMzkwMjA1MH0.86owJoyHXTF5tikrFoQpDuDA-UJve_GWcq7qAvKBcn8")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("http do err", err)
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

func printUsageErrorAndExit(message string) {
	fmt.Fprintln(os.Stderr, "ERROR:", message)
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Available command line options:")
	flag.PrintDefaults()
	os.Exit(64)
}

func printErrorAndExit(code int, format string, values ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", fmt.Sprintf(format, values...))
	fmt.Fprintln(os.Stderr)
	os.Exit(code)
}
