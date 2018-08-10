package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
	"fmt"
	"time"
	"crypto/sha1"
	"encoding/hex"
)

type Xuser struct {
	Id   int64  `"xorm:autoincr"`
	Name string `"xorm:varchar(255)"`
	Sex  int64  `"xorm:int(11)"`
	Age  int64  `"xorm:int(11)"`
}

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("mysql", "dc_datacenter:dc2014local@(192.168.7.21:3306)/field_case_test?charset=utf8")
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
}

func main() {
	//err := engine.CreateTables(new(Xuser))
	//if err != nil {
	//	panic("err")
	//}
	fmt.Println("StartTime:" + time.Now().Format("2006-01-02 15:04:05"))
	for a := 0; a <= 1000; a++ {
		xuser := new(Xuser)
		//sql := "INSERT INTO User(Name, Sex, Age) values(SHA1('zhangsan'),1,1)"
		hash := sha1.New()
		hash.Write([]byte("zhangsan"))
		name := hex.EncodeToString(hash.Sum(nil))
		xuser.Name = name
		xuser.Age = 1
		xuser.Sex = 1
		_, err := engine.Insert(xuser)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("EndTime:" + time.Now().Format("2006-01-02 15:04:05"))
}
