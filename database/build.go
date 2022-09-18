package database

import (
	"database/sql"
	"fmt"
	"log"
)

func BuildFromSQL() {
	postgres := DB.Postgres
	buildUsers(postgres)
}

func buildUsers(postgres *sql.DB) {
	_, err := postgres.Exec(`CREATE SCHEMA IF NOT EXISTS "Users";`)
	checkBuildError(err, postgres, "Users Schema")
	_, err = postgres.Exec(`
		CREATE TABLE IF NOT EXISTS "Users"."Users"
		(
			"Username" character varying(30) COLLATE pg_catalog."default" NOT NULL,
			"Email" character varying(30) COLLATE pg_catalog."default",
			"Access" smallint default 0,
			"CountryCode" integer,
			"PasswordHash" character varying(30) COLLATE pg_catalog."default" NOT NULL,
			"SessionID" integer,
			"CreatedAt" integer,
			"LastLogin" integer,
			"ID" integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
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
