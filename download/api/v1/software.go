package v1

import (
	"deepgo/download/models"
	"deepgo/download/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strings"
)

//获取软件列表
func GetSoftware(c *gin.Context) {
	data := make(map[string]interface{})
	var software models.Software
	command := exec.Command("ls", "/Users/xiaowei/go/src/deepgo")
	bytes, err := command.Output()
	if err != nil {
		fmt.Println("err", err)
	}
	result := string(bytes)
	tmp := strings.Split(strings.TrimSpace(result), "\n")
	var softwareList []models.Software
	count := 1
	fmt.Println(c.ClientIP())
	for _, s := range tmp {
		software.Id = count
		software.Name = s
		url := fmt.Sprintf("<a href='http://%s:8002/%s'>download</a>", setting.URL, s)
		software.Url = url
		softwareList = append(softwareList, software)
		count += 1
	}
	//fmt.Println(softwareList)

	if err != nil {
		fmt.Println(err)
	}
	//code := e.SUCCESS
	data["rows"] = softwareList
	data["total"] = len(tmp)
	c.JSON(http.StatusOK, gin.H{
		//	"code" : code,
		//	"msg" : e.GetMsg(code),
		//	"data" : data,
		"rows":  softwareList,
		"total": len(tmp),
	})

}

////新增文章标签
//func AddTag(c *gin.Context) {
//}
//
////修改文章标签
//func EditTag(c *gin.Context) {
//}
//
////删除文章标签
//func DeleteTag(c *gin.Context) {
//}
