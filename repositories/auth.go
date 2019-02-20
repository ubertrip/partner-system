package repositories

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func GetUserByLogin(login, password string) (bool, int) {

	hash := ""
	ID := 0

	err := Get().QueryRow("SELECT login, password, ID FROM `users` WHERE login=?", login).Scan(
		&login,
		&hash,
		&ID)

	fmt.Println(login, password, ID, hash)

	if err != nil {
		fmt.Println(err, "error")

		return false, 0
	}
	return CheckPassword(password, hash), ID

}

func GetStatus() bool {
	return false
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

func CreateUser(login, password string) (err error) {
	_, err = Get().Exec("INSERT INTO users (login, password) VALUES (?, ?)", login, password)
	return
}
