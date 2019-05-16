package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(192.168.2.189:3306)/deepvideo?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	user := User{
		Name: "slene",
	}
	id, err := o.Insert(&user)
	fmt.Printf("ID:%d,ERR:%v\n", id, err)

	user.Name = "xiaowei"
	num, err := o.Update(&user)
	fmt.Printf("NUM:%d,ERR:%v\n", num, err)

	u := User{Id: user.Id}
	o.Read(&u)
	fmt.Println(u)

	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
