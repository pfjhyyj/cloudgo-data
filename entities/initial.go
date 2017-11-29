package entities

import (
	//  use mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	//https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	myEngine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true")
	checkErr(err)
	err = myEngine.Sync2(new(UserInfo))
	checkErr(err)
	engine = myEngine
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
