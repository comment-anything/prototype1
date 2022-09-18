package database

import (
	"fmt"
	"log"
	"time"
)

type User struct {
	ID           int
	Username     string
	CreatedAt    time.Time
	LastLogin    time.Time
	Email        string
	Access       string
	CountryCode  int
	PasswordHash string
	SessionID    int
}

// CreateUser creates a User struct and adds it to the database. It returns that struct.
func CreateUser(username string, email string, access string, countryCode int, password string) (*User, error) {

	user := User{
		/** We need to validate the uniqueness of usernames, unless we will allow changing of usernames and include an id field. We will steel to validate to ensure they don't have sql breaking characters like commas and quotes. */
		Username:  username,
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
		/** We need to validate the email. */
		Email:       email,
		Access:      "Poster",
		CountryCode: countryCode,
		/** We need to encrypt + salt the password for secure storage. */
		PasswordHash: password,
	}

	fmt.Printf("User creation time type %t", user.CreatedAt)

	insertStatement := fmt.Sprintf(`
		Insert into "Users"."Users"(
		"Username",
		"CreatedAt",
		"LastLogin",
		"Email",
		"CountryCode",
		"PasswordHash"
		) values('%s', %v, %v, '%s', %d, '%s');`,
		user.Username,
		user.CreatedAt.Unix(),
		user.LastLogin.Unix(),
		user.Email,
		user.CountryCode,
		user.PasswordHash)

	postgres := DBConnector.Connect()

	_, err := postgres.Exec(insertStatement)
	if err != nil {
		postgres.Close()
		log.Fatalf(err.Error())
	}

	postgres.Close()

	return &user, nil
}
