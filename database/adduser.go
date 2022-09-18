package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type User struct {
	ID           int
	Username     string
	CreatedAt    int64
	LastLogin    int64
	Email        string
	Access       UserAccessLevel
	CountryCode  int
	PasswordHash string
	SessionID    sql.NullInt64
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
		CreatedAt: time.Now().Unix(),
		LastLogin: time.Now().Unix(),
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
		user.CreatedAt,
		user.LastLogin,
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

func GetUser(id int) (*User, error) {
	var user User
	queryStatement := fmt.Sprintf(`
		select "ID", "Username", "Email", "Access", "CountryCode", "PasswordHash", "CreatedAt", "LastLogin", "SessionID" from "Users"."Users" where "ID"=%v`, id)
	rows, err := DB.Postgres.Query(queryStatement)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.ID,
			&user.Username,
			&user.Email,
			&user.Access,
			&user.CountryCode,
			&user.PasswordHash,
			&user.CreatedAt,
			&user.LastLogin,
			&user.SessionID)
		if err != nil {
			return nil, err
		} else {
			return &user, err
		}
	}
	return nil, errors.New("Unable to find user.")
}
