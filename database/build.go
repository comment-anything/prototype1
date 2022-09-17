package database

import (
	"database/sql"
	"fmt"
	"log"
)

func BuildFromSQL() {
	postgres := DBConnector.Connect()
	buildUsers(postgres)
	postgres.Close()
}

func buildUsers(postgres *sql.DB) {
	_, err := postgres.Exec(`CREATE SCHEMA IF NOT EXISTS "Users";`)
	checkBuildError(err, postgres, "Users Schema")
	_, err = postgres.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname='AccessLevels')
			THEN
				CREATE TYPE "Users"."AccessLevels" AS ENUM (
					'Poster',
					'DomainModerator',
					'GlobalModerator',
					'Administrator'
				);
			END IF;
		END$$;`)
	checkBuildError(err, postgres, "Users.AccessLevels Enum")
	_, err = postgres.Exec(`
		CREATE TABLE IF NOT EXISTS "Users"."Users"
		(
			"ID" integer NOT NULL,
			"Username" varchar(30) NOT NULL,
			"CreatedAt" timestamp,
			"LastLogin" timestamp,
			"Email" varchar(30),
			"Access" "Users"."AccessLevels",
			"CountryCode" integer,
			"PasswordHash" varchar(30) NOT NULL,
			"SessionID" integer,
			CONSTRAINT "Users_pkey" PRIMARY KEY ("ID")
		);
	`)
	checkBuildError(err, postgres, "Users Table")
}

func checkBuildError(err error, db *sql.DB, error_on string) {
	if err != nil {
		db.Close()
		fmt.Printf("Problem in the the build process on %s.\n", error_on)
		log.Fatalf(err.Error())
	}
}
