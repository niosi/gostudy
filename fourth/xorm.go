package main
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

type User struct {
	Id int64 "xorm:autoincr"
	Name string
	Sex int
	Age int
}

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("mysql", "root:root@/world?charset=utf8")
	f, err := os.Create("sql.log")
	if err != nil {
		println(err.Error())
		return
	}
	engine.SetLogger(xorm.NewSimpleLogger(f))
	engine.ShowSQL(true)
	engine.ShowExecTime(true);
}

func main() {
	err := engine.CreateTables(new(User))
	if err!=nil {
		panic("err")
	}
	var sql = "INSERT INTO User(Name, Sex, Age) values('zhangsan',1,1)"
	result,err := engine.Exec(sql)
	if err == nil {
		id,err := result.RowsAffected()
		if err!=nil {
			panic("err")
		}
		println(id)
	}

	user := &User{Id:3}
	has, err := engine.Get(user)
	if has {
		println(user.Name)
	}
}
