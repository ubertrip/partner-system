package repositories

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync/atomic"
	"os"
	"fmt"
	configuration "github.com/ubertrip/partner-system/config"
)

var (
	atomicDB atomic.Value
)

func InitDB() {
	var db *sql.DB
	var err error

	onError := func(err error) {
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
			return
		}
	}

	config := configuration.Get()

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Schema)

	db, err = sql.Open("mysql", connection)
	onError(err)

	err = db.Ping()
	onError(err)

	db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(128)

	atomicDB.Store(db)

	fmt.Println("Database OK")
}

func Get() *sql.DB {
	return atomicDB.Load().(*sql.DB)
}
