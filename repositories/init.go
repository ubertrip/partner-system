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

func GetUserByLogin(login, password string) bool {
	// a := login 
	// b := password
	err := Get().QueryRow("SELECT login, password FROM `users` WHERE login=? and password=?", login, password).Scan(		
		&login,
		&password)
	// err := Get().QueryRow("SELECT password FROM users WHERE name=?", login).Scan(&password)
	// fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		
		return false
	}
	
	return true
}

// func Login() {
// 	fmt.Println(">>", GetUserByLogin("123", "123"))
// 	}

func Get() *sql.DB {
	return atomicDB.Load().(*sql.DB)
}

// hc := http.Client{}
// req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))

//     form := url.Values{}
//     form.Add("ln", c.ln)
//     form.Add("ip", c.ip)
//     form.Add("ua", c.ua)
//     req.PostForm = form
//     req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

//     glog.Info("form was %v", form)
//     resp, err := hc.Do(req)
