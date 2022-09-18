package database

import (
	"fmt"
	"time"
)

type User struct {
	ID           int
	Username     string
	CreatedAt    time.Time
	LastLogin    time.Time
	Email        string
	Access       UserAccessLevel
	CountryCode  int
	PasswordHash string
	SessionID    int
}

type UserAccessLevel byte

const (
	UALPoster          UserAccessLevel = 0
	UALDomainModerator UserAccessLevel = 1
	UALGlobalModerator UserAccessLevel = 2
	UALAdministrator   UserAccessLevel = 3
)

// CreateUser creates a User struct and adds it to the database. It returns that struct.
func CreateUser(username string, email string, access UserAccessLevel, countryCode int, password string) (*User, error) {

	user := User{
		/** We need to validate the uniqueness of usernames, unless we will allow changing of usernames and include an id field. We will steel to validate to ensure they don't have sql breaking characters like commas and quotes. */
		Username:  username,
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
		/** We need to validate the email. */
		Email:       email,
		Access:      access,
		CountryCode: countryCode,
		/** We need to encrypt + salt the password for secure storage. */
		PasswordHash: password,
	}

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

	postgres := DB.Postgres

	_, err := postgres.Exec(insertStatement)
	if err != nil {
		fmt.Println("User insert err: ", err.Error())
	}

	return &user, nil
}
