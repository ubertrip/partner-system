package repositories

import (
	"database/sql"
	"fmt"
	"os"
	"sync/atomic"

	_ "github.com/go-sql-driver/mysql"
	configuration "github.com/ubertrip/partner-system/config"
	"golang.org/x/crypto/bcrypt"
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

func GetUserByLogin(login, password string) (bool, int) {

	hash := ""
	ID := 0

	err := Get().QueryRow("SELECT login, password, ID FROM `users` WHERE login=?", login).Scan(
		&login,
		&hash,
		&ID)

	fmt.Println(password, hash)
	fmt.Println(login, password, hash)
	fmt.Println(login, password, ID)

	if err != nil {
		fmt.Println(err, "error")
		fmt.Println(password, hash)

		return false, 0
	}
	fmt.Println(password, hash)
	return CheckPassword(password, hash), ID

}

func GetStatus() bool {
	return false
}

func Get() *sql.DB {
	return atomicDB.Load().(*sql.DB)
}

const (
	defaultCost = 12
)

func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func HashPassword(password string) string {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), defaultCost); err == nil {
		return string(hashedPassword)
	}
	return ""
}

// func init() {
// 	ID := 1
// 	pass := GetUserByLogin("123", "123", (ID))
// 	fmt.Println(pass)
// }
