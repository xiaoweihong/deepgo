package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os/exec"
	"strings"
	"time"
)

type Programe struct {
	Id         int
	Name       string `orm:"unique"`
	CreateTime time.Time
}
type Arcee struct {
	Programe
}

type Crusader struct {
	Programe
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.2.189:3306)/deepvideo?charset=utf8", 30)
	orm.RegisterModel(new(Arcee))
	orm.RegisterModel(new(Crusader))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RunSyncdb("default", false, true)

}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	insertObject(o, Arcee{})
	insertObject(o, Crusader{})

}

func insertObject(o orm.Ormer, programe interface{}) {

	switch v := programe.(type) {
	case Arcee:
		var arcee Arcee
		arcee = v
		var objectList []Arcee
		resultList := getVersionList("arcee")
		for _, v := range resultList {
			resultTmp := strings.Split(v, " ")
			arcee.Name = resultTmp[0]
			arcee.CreateTime = parseTime(resultTmp[1])
			objectList = append(objectList, arcee)
			//result, err := o.Insert(&arcee)
		}
		_, err := o.InsertMulti(100, objectList)
		if err != nil {
			fmt.Println(err)
		}
	case Crusader:
		var crusader Crusader
		crusader = v
		var objectList []Crusader
		resultList := getVersionList("crusader")
		for _, v := range resultList {
			resultTmp := strings.Split(v, " ")
			crusader.Name = resultTmp[0]
			crusader.CreateTime = parseTime(resultTmp[1])
			objectList = append(objectList, crusader)
		}
		_, err := o.InsertMulti(100, objectList)
		if err != nil {
			fmt.Println(err)
		}
		//case Arcee:
		//	var arcee Arcee
		//	arcee=v
		//	var objectList []Arcee
		//	resultList := getVersionList("arcee")
		//	for _,v := range resultList{
		//		resultTmp:=strings.Split(v," ")
		//		arcee.Name=resultTmp[0]
		//		arcee.CreateTime=parseTime(resultTmp[1])
		//		objectList=append(objectList,arcee)
		//		//result, err := o.Insert(&arcee)
		//	}
		//	_, err := o.InsertMulti(100, objectList)
		//	if err != nil {
		//		fmt.Println(err)
		//	}
	}

	//fmt.Println(objectList)

}

//func saveObjectToDb(resultList []string,object interface{})  {
//	for _,v := range resultList{
//		resultTmp:=strings.Split(v," ")
//		object.Name=resultTmp[0]
//		object.CreateTime=parseTime(resultTmp[1])
//		objectList=append(objectList,arcee)
//		//result, err := o.Insert(&arcee)
//	}
//	_, err := o.InsertMulti(100, objectList)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

func getVersionList(name string) (resultList []string) {
	command := exec.Command("bash", "getVersion.sh", name)
	bytes, err := command.Output()
	if err != nil {
		fmt.Println("err", err)
	}
	result := string(bytes)
	resultList = strings.Split(strings.Trim(result, "\n"), "\n")
	return
}

func parseTime(time_a string) (time_s time.Time) {
	string_tmp := strings.Split(time_a, "-")
	timeLayout := "2006-01-02"
	year := string_tmp[2]
	month := string_tmp[1]
	day := string_tmp[0]
	switch month {
	case "Jan":
		month = "01"
	case "Feb":
		month = "02"
	case "Mar":
		month = "03"
	case "Apr":
		month = "04"
	case "May":
		month = "05"
	case "Jun":
		month = "06"
	case "Jul":
		month = "07"
	case "Aug":
		month = "08"
	case "Sep":
		month = "09"
	case "Oct":
		month = "10"
	case "Nov":
		month = "11"
	case "Dec":
		month = "12"
	}
	parse_string := fmt.Sprintf("%s-%s-%s", year, month, day)
	time_s, err := time.Parse(timeLayout, parse_string)
	if err != nil {
		fmt.Println(err)
	}
	return
}
