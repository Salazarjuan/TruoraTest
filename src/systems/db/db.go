package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

/*func Conect() (db *xorm.Engine, err error) {

	db, err = xorm.NewEngine("mysql", "test:test@tcp(localhost:3306)/test?charset=utf8")

	if err != nil {
		return
	}
	db.DBMetas()

	return

}*/

func Connect(host string, port string, user string, pass string, database string, options string) (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&"+options)
}
