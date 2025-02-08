package user

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
}

func NewUser(db *sql.DB, u User) error {

	sqlStatement := "insert into users(username, password, email) values ($1, $2, $3)"
	_, err := db.Exec(sqlStatement, u.Username, u.Password, u.Email)
	if err != nil {
		fmt.Println("Insert Error: %w", err)
		return err
	}

	log.Printf("New User: %s[%s]\n", u.Username, u.Email)
	return nil
}

func CheckUser(db *sql.DB, email string) error {
	sqlStatement := "select * from users where email = $1"
	_, err := db.Exec(sqlStatement, email)
	if err != nil {
		fmt.Println("Insert Error: %w", err)
		return err
	}

	return nil

}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	sqlStatement := "select * from users where email = $1"
	var u User
	err := db.QueryRow(sqlStatement, email).Scan(&u.Id, &u.Username, &u.Email, &u.Password)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user with email %s not found", email)
	}
	if err != nil {
		return User{}, fmt.Errorf("cannot get user with email %s: %w", email, err)
	}

	return u, nil
}
